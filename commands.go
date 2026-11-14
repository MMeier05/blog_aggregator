package main

import (
	"errors"
)

type command struct {
	name string
	args []string
}

type commands struct {
	functions map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.functions[cmd.name]
	if ok {
		err := value(s, cmd)
		if err != nil {
			return err
		}
	} else {
		return errors.New("Command not found!\n")
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.functions[name] = f
}



