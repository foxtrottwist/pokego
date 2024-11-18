package commands

import (
	"errors"
	"fmt"
)

func exploreCommand(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("a location area name must be provided")
	}

	la, err := c.client.LocationArea(args[0])
	if err != nil {
		return err
	}

	c.write(fmt.Sprintf("Exploring %s...\n", args[0]))
	c.write("Found Pokemon: \n")
	for _, p := range la.PokemonEncounters {
		c.write(fmt.Sprintf("- %s\n", p.Pokemon.Name))
	}

	c.write("\n")
	return nil
}
