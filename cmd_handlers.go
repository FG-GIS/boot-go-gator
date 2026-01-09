package main

import (
	"errors"
	"log"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("Error, not enough arguments, Username is required for login.")
	}
	if len(cmd.args) > 1 {
		return errors.New("Error, too many arguments.")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	log.Println("GATOR -- User correctly set.")
	return nil
}
