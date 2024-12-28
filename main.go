package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Bayan2019/rss/internal/config"

	_ "github.com/lib/pq"
)

// Create a state struct that holds a pointer to a config
type state struct {
	cfg *config.Config
}

func main() {
	// read the config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	// store the config in a new instance of the state struct.
	programState := &state{
		cfg: &cfg,
	}

	// Create a new instance of the commands struct
	// with an initialized map of handler functions.
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	// egister a handler function for the login command.
	cmds.register("login", handlerLogin)

	// Use os.Args to get the command-line arguments passed in by the user.
	if len(os.Args) < 2 {
		// If there are fewer than 2 arguments,
		// print an error message to the terminal and exit.
		fmt.Println("Usage: cli <command> [args...]")
		log.Fatal(fmt.Errorf("Not enough arguments"))
		return
	}

	// to split the command-line arguments
	// into the command name
	cmdName := os.Args[1]
	// and the arguments slice to create a command instance.
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
