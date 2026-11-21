package main

import (
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	registry map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	curr_cmd, ok := c.registry[cmd.name]
	if !ok {
		return fmt.Errorf("The %s command does not exist", cmd.name)
	}

	return curr_cmd(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registry[name] = f
}
