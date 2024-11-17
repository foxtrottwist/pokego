package commands

import (
	"errors"
	"fmt"
)

func exploreCommand(c *config, args ...string) error {
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
