package commands

import (
	"errors"
	"fmt"
)

func inspectCommand(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("a Pokémon name must be provided")
	}

	p, ok := c.pokedex[args[0]]
	if !ok {
		return errors.New("You haven't caught that Pokémon")
	}

	c.print(fmt.Sprintf("Name: %s\n", p.Name))
	c.print(fmt.Sprintf("Height: %d\n", p.Height))
	c.print(fmt.Sprintf("Weight: %d\n", p.Weight))

	c.print("Stats: \n")
	for _, info := range p.Stats {
		c.print(fmt.Sprintf("- %s: "+"%d\n", info.Stat.Name, info.BaseStat))
	}

	c.print("Types: \n")
	for _, info := range p.Types {
		c.print(fmt.Sprintf("- %s\n", info.Type.Name))
	}

	c.print("\n")
	return nil
}
