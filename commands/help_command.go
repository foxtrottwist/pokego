package commands

import "fmt"

func helpCommand(c *config, args ...string) error {
	c.write("\nWelcome to the PokéGo Pokédex!\nUsage:\n\n")

	for _, cmd := range Commands() {
		c.write(fmt.Sprintf("%s: %s\n", cmd.name, cmd.description))
	}

	c.write("\n")
	return nil
}
