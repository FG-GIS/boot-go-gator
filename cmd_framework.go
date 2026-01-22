package main

import (
	"errors"
	"slices"
)

type logging struct {
	flag bool
	path string
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commandList[cmd.name]
	if !ok {
		return errors.New("GATOR -- Error, command not mapped.")
	}
	err := f(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
}

func (c *command) validate() (bool, logging, error) {
	verbose := false
	log := logging{
		flag: false,
		path: "",
	}
	if slices.Contains(c.args, "-v") {
		verbose = true
		c.args = slices.DeleteFunc(c.args, func(s string) bool {
			return s == "-v"
		})
	}
	if slices.Contains(c.args, "-log") {
		log.flag = true
		idx := slices.Index(c.args, "-log")
		log.path = c.args[idx+1]
		c.args = slices.Delete(c.args, idx, idx+1)
	}
	cmdLen := len(c.args)
	switch c.name {
	case "login", "register":
		if cmdLen < 1 {
			return verbose, log, errors.New("GATOR -- Error, not enough arguments, Username is required.")
		}
		return verbose, log, nil
	default:
		return verbose, log, errors.New("GATOR -- Error validating command, command name not found.")
	}
}
