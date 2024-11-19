package commands

import "fmt"

func helpCommand(c *config, args ...string) error {
	c.print("\nWelcome to the PokéGo Pokédex!\nUsage:\n\n")

	for _, cmd := range Commands() {
		c.print(fmt.Sprintf("%s: %s\n", cmd.name, cmd.description))
	}

	c.print("\n")
	return nil
}
