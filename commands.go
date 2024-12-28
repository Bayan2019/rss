package main

import "errors"

// Create a command struct.
type command struct {
	// A command contains a name
	Name string
	// and a slice of string arguments.
	Args []string
}

// Create a commands struct.
// This will hold all the commands the CLI can handle.
type commands struct {
	// Add a map[string]func(*state, command) error field to it.
	// This will be a map of command names to their handler functions.
	registeredCommands map[string]func(*state, command) error
}

// This method registers a new handler function for a command name.
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// This method runs a given command with the provided state if it exists.
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
