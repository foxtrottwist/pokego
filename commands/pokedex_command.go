package commands

import (
	"errors"
	"fmt"
)

func pokedexCommand(c *config, _ ...string) error {
	if len(c.pokedex) == 0 {
		return errors.New("Pokédex is empty")
	}
	for name := range c.pokedex {
		c.print(fmt.Sprintf(" - %s\n", name))
	}
	return nil
}
