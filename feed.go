package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/PenguinKG/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAggregator(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Error: Missing Arguments (Request span required)")
	}
	time_between_reqs, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}
	ticker := time.NewTicker(time_between_reqs)

	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("Error: %w", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	fetchedParams := database.MarkFeedFetchedParams{
		ID:            feed.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	err = s.db.MarkFeedFetched(context.Background(), fetchedParams)
	if err != nil {
		return err
	}

	fetch, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}
	for _, item := range fetch.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Error: Missing Arguments (Name and URL required)")
	}

	currTime := time.Now()

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: currTime,
		UpdatedAt: currTime,
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: currTime,
		UpdatedAt: currTime,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	fmt.Printf("%v created new feed: %s (%v) [%v]\n", feed.UserID, feed.Name, feed.Url, feed.CreatedAt)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	for _, feed := range feeds {
		username, err := s.db.GetUsernameByID(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("Error: %w", err)
		}
		fmt.Printf("* %s (%s) [Created by: %s]\n", feed.Name, feed.Url, username)
	}

	return nil
}
