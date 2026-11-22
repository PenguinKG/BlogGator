package main

import (
	"context"
	"fmt"
	"time"

	"github.com/PenguinKG/gator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("Error: Missing Argument (URL required)")
	}

	currTime := time.Now()
	feed, err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.Username)
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

func handlerFollowing(s *state, cmd command) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.Username)
	if err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	for _, follow := range feedFollows {
		fmt.Printf("* %s follows %s [Started following on: %v]\n", follow.Username, follow.FeedName, follow.CreatedAt)
	}

	return nil
}
