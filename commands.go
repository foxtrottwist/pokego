package main

import (
	"fmt"
	"os"

	fetch "github.com/foxtrottwist/pokego/fetch"
)

type command struct {
	name        string
	description string
	run         func() error
}

func commands() map[string]command {
	return map[string]command{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			run:         exitCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			run:         helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas",
			run:         mapCommand,
		},
	}
}

func exitCommand() error {
	os.Exit(0)
	return nil
}

func helpCommand() error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}

func mapCommand() error {
	loc, err := fetch.LocationAreas()
	if err != nil {
		return err
	}

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	return nil
}
