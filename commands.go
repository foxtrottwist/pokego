package main

import (
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	run         func() error
}

func commands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			run:         helpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			run:         exitCommand,
		},
	}
}

func helpCommand() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}
