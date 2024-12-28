package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Bayan2019/rss/internal/database"
	"github.com/google/uuid"
)

// Create a register handler function
func handlerRegister(s *state, cmd command) error {
	// Ensure that a name was passed in the args.
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(
		// Pass context.Background() to the query to create an empty Context argument.
		context.Background(),
		database.CreateUserParams{
			// Use the uuid.New() function to generate a new UUID for the user.
			ID: uuid.New(),
			// created_at and updated_at should be the current time.
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			// Use the provided name.
			Name: name,
		},
	)
	if err != nil {
		return err
	}

	// Set the current user in the config to the given name.
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	// Print a message to the terminal that the user has been set.
	fmt.Println("User created successfully!")
	printUser(user)
	return nil
}

// Create a login handler function
func handlerLogin(s *state, cmd command) error {
	// If the command's arg's slice is empty, return an error;
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	// Update the login command handler to error (and exit with code 1)
	// if the given username doesn't exist in the database.
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	// Use the state's access to the config struct
	// to set the user to the given username.
	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	// Print a message to the terminal that the user has been set.
	fmt.Println("User switched successfully!")
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}
