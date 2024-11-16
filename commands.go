package main

import (
	"errors"
	"fmt"
	"os"
)

type command struct {
	name        string
	description string
	run         func(*config, ...string) error
}

func commands() map[string]command {
	return map[string]command{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			run:         exitCommand,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of Pokemon found in the location area",
			run:         exploreCommand,
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

func exitCommand(*config, ...string) error {
	os.Exit(0)
	return nil
}

func exploreCommand(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("a location area name must be provided")
	}

	la, err := c.client.LocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon: ")
	for _, p := range la.PokemonEncounters {
		fmt.Printf("- %s\n", p.Pokemon.Name)
	}

	fmt.Println()
	return nil
}

func helpCommand(*config, ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}

func mapCommand(c *config, args ...string) error {
	la, err := c.client.LocationAreas(c.next)
	if err != nil {
		return err
	}

	c.next = la.Next
	c.previous = la.Previous

	for _, area := range la.Results {
		fmt.Println(area.Name)
	}

	fmt.Println()
	return nil
}

func mapbCommand(c *config, args ...string) error {
	if c.previous == nil {
		return errors.New("cannot go back, you're on the first page")
	}

	la, err := c.client.LocationAreas(c.previous)
	if err != nil {
		return err
	}

	c.next = la.Next
	c.previous = la.Previous

	for _, area := range la.Results {
		fmt.Println(area.Name)
	}

	fmt.Println()
	return nil
}
