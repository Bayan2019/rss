package main

import "fmt"

// Create a login handler function
func handlerLogin(s *state, cmd command) error {
	// If the command's arg's slice is empty, return an error;
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	// Use the state's access to the config struct
	// to set the user to the given username.
	err := s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	// Print a message to the terminal that the user has been set.
	fmt.Println("User switched successfully!")
	return nil
}
