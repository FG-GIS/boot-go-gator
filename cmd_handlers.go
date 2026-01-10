package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/FG-GIS/boot-go-gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("GATOR -- Error, not enough arguments, Username is required for login.")
	}
	if len(cmd.args) > 1 {
		return errors.New("GATOR -- Error, too many arguments.")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err != nil {
		return fmt.Errorf("GATOR -- Error, non-existant username.\n %v", err)
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("GATOR -- User correctly set.")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Error, not enough arguments, Username is required for registration.")
	}
	if len(cmd.args) > 1 {
		return errors.New("Error, too many arguments.")
	}
	if _, err := s.db.GetUser(context.Background(), cmd.args[0]); err == nil {
		return errors.New("User already registered, exiting.")
	}
	usr, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		return err
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("GATOR -- User correctly registered.")
	fmt.Printf("User generated: %v", usr)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("Error, too many arguments.")
	}
	err := s.db.Reset(context.Background())
	if err != nil {
		fmt.Println("Error re-setting table")
		return err
	}
	fmt.Println("GATOR -- Table cleansed")
	return nil
}

func handlerGetUsers(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("Error, too many arguments.")
	}
	usrSlice, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("Error gathering users from table")
		return err
	}
	for _, usr := range usrSlice {
		if usr == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", usr)
		} else {
			fmt.Printf("* %v\n", usr)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	// if len(cmd.args) == 0 {
	// 	return errors.New("Error, not enough arguments, url required.")
	// }
	if len(cmd.args) > 1 {
		return errors.New("Error, too many arguments.")
	}
	if len(cmd.args) == 0 {
		cmd.args = append(cmd.args, "https://www.wagslane.dev/index.xml")
	}
	feed, err := fetchFeed(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Println("GATOR -- Error fetching feed.")
		return err
	}
	fmt.Println("GATOR -- fetched feed ==> ", feed)
	return nil
}

func handlerFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return errors.New("GATOR -- Error, not enough arguments, name and url required.")
	}
	currentUsr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		fmt.Println("GATOR -- Error getting current user.")
		return err
	}
	feedEntry := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    currentUsr.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedEntry)
	if err != nil {
		fmt.Println("GATOR -- Error inserting feed entry.")
		return err
	}
	fmt.Println("GATOR -- Feed inserted successfully.")
	fmt.Println(feed)
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return errors.New("GATOR -- Error, too many arguments.")
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("GATOR -- Error, retrieving feed entries.")
		return err
	}
	if len(feeds) == 0 {
		fmt.Println("GATOR -- Error, no feeds to retrieve.")
		return nil
	}
	printFeeds(feeds)
	return nil
}
