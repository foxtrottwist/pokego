package commands

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/foxtrottwist/pokego/client"
)

type Config struct {
	client.Client
	next     *string
	previous *string
}

type command struct {
	name        string
	description string
	Run         func(*Config, ...string) error
}

func Commands() map[string]command {
	return map[string]command{
		"cache": {
			name:        "cache",
			description: "Manipulates the PokéGo cache",
			Run:         cacheCommand,
		},
		"clear": {
			name:        "clear",
			description: "Clears the PokéGo output",
			Run:         clearCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits PokéGo",
			Run:         exitCommand,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of Pokemon found in the location area",
			Run:         exploreCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Run:         helpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			Run:         mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			Run:         mapbCommand,
		},
	}
}

const (
	CLEAN                  = "clean"
	LS                     = "ls"
	NO_CACHE_COMMAND_ERROR = "no `cache` command provided\nManipulates the PokéGo cache\n\nUsage: \n\ncache clean: removes all items from the cache\ncache ls: lists all items in the cache"
)

func cacheCommand(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New(NO_CACHE_COMMAND_ERROR)
	}

	switch args[0] {
	case CLEAN:
		c.Client.CleanCache()
	case LS:
		c.Client.ListCache()
	default:
		fmt.Printf("cache %s: unknown command\n", args[0])
	}

	return nil
}

func clearCommand(*Config, ...string) error {
	if runtime.GOOS == "windows" {
		return clearHelper(exec.Command("cmd", "/c", "cls"))
	}
	return clearHelper(exec.Command("clear"))
}

func clearHelper(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func exitCommand(*Config, ...string) error {
	os.Exit(0)
	return nil
}

func exploreCommand(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("a location area name must be provided")
	}

	la, err := c.Client.LocationArea(args[0])
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

func helpCommand(*Config, ...string) error {
	fmt.Printf("\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, cmd := range Commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}

func mapCommand(c *Config, args ...string) error {
	la, err := c.Client.LocationAreas(c.next)
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

func mapbCommand(c *Config, args ...string) error {
	if c.previous == nil {
		return errors.New("cannot go back, you're on the first page")
	}

	la, err := c.Client.LocationAreas(c.previous)
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
