package main

import (
	"errors"
	"fmt"
	"github.com/MMeier05/blog_aggregator/internal/database"
	"github.com/google/uuid"
	"time"
	"context"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return errors.New("usage: <url>")
	}
	user, err := s.db.GetUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	feed, err := s.db.GetFeed(context.Background(), cmd.args[0])
	params := database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID: user.ID,
		FeedID: feed.ID,
	}
	inserted_feed_follow, err:= s.db.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return err
	}
	fmt.Printf("Created feed: %s (%s)\n", inserted_feed_follow.FeedName, inserted_feed_follow.UserName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	followedFeeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.config.CurrentUserName)
	if err != nil {
		return err
	}
	fmt.Printf("%s's Feeds:\n", s.config.CurrentUserName)
	for _, followedFeed := range followedFeeds {
		fmt.Printf("* %s\n", followedFeed.FeedName)
	}
	return nil
}
