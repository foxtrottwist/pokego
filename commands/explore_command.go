package commands

import (
	"errors"
	"fmt"
)

func exploreCommand(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("a location area name must be provided")
	}

	la, err := c.client.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	c.print(fmt.Sprintf("Exploring %s...\n", args[0]))
	c.print("Found Pokemon: \n")
	for _, p := range la.PokemonEncounters {
		c.print(fmt.Sprintf("- %s\n", p.Pokemon.Name))
	}

	c.print("\n")
	return nil
}
