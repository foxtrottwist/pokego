package main

import (
	"fmt"
	"os"

	fetch "github.com/foxtrottwist/pokego/fetch"
)

type config struct {
	Next     string
	Previous string
}

type command struct {
	name        string
	description string
	run         func(*config) error
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
			description: "Displays the names of the next 20 location areas",
			run:         mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			run:         mapbCommand,
		},
	}
}

func exitCommand(*config) error {
	os.Exit(0)
	return nil
}

func helpCommand(*config) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}

func mapCommand(c *config) error {
	la, err := fetch.LocationAreas(c.Next)
	if err != nil {
		return err
	}

	c.Next = la.Next
	c.Previous = la.Previous

	for _, area := range la.Results {
		fmt.Println(area.Name)
	}

	fmt.Println()
	return nil
}

func mapbCommand(c *config) error {
	la, err := fetch.LocationAreas(c.Previous)
	if err != nil {
		return err
	}

	c.Next = la.Next
	c.Previous = la.Previous

	for _, area := range la.Results {
		fmt.Println(area.Name)
	}

	fmt.Println()
	return nil
}
