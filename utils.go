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
