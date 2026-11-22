package main

import (
	"context"
	"fmt"
	"time"

	"github.com/PenguinKG/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Error: Missing Argument (URL required)")
	}

	currTime := time.Now()
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
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

	follow, err := s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	fmt.Printf("%s is now following %s [Followed on: %v]\n", follow.Username, follow.FeedName, follow.CreatedAt)
	return nil
}

func handlerFollowing(s *state, cmd command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	for _, follow := range feedFollows {
		fmt.Printf("* %s follows %s [Started following on: %v]\n", follow.Username, follow.FeedName, follow.CreatedAt)
	}

	return nil
}

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Error: Missing Argument (URL required)")
	}

	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	unfollowParams := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFeedFollow(context.Background(), unfollowParams)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	fmt.Printf("%s is no longer following %s\n", user.Name, feed.Name)
	return nil
}
