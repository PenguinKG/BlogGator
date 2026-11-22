package main

import (
	"context"
	"fmt"
	"time"

	"github.com/PenguinKG/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAggregator(s *state, cmd command) error {
	fetch, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}
	fmt.Printf("Feed: %v\n", fetch)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("Error: Missing Arguments (Name and URL required)")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.Username)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
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
