package main

import (
	"fmt"

	"github.com/FG-GIS/boot-go-gator/internal/database"
)

func printFeeds(feedSlice []database.GetFeedsRow) {
	fmt.Printf("GATOR -- Printing out feeds (%d)", len(feedSlice))
	for _, f := range feedSlice {
		fmt.Printf("* ID:							%s\n", f.ID)
		fmt.Printf("* Name:							%s\n", f.Name)
		fmt.Printf("* URL:							%s\n", f.Url)
		fmt.Printf("* User:							%s\n", f.User)
	}
}

func printFollowing(followSlice []database.GetFeedFollowsForUserRow, user string) {
	fmt.Printf("GATOR -- User ==> %s - is following:\n", user)
	for _, follow := range followSlice {
		fmt.Printf("* %s\n", follow.FeedName)
	}
}
