package commands

import (
	"errors"
	"fmt"
)

const (
	clean = "clean"
	ls    = "ls"
)

func cacheCommand(c *config, args ...string) error {
	descriptor := fmt.Sprintf("Manipulates the PokéGo cache\n\nUsage: \n\ncache %s: removes all items from the cache\ncache %s: lists all items in the cache", clean, ls)

	if len(args) == 0 {
		return errors.New(fmt.Sprint("no `cache` command provided\n\n", descriptor))
	}

	switch args[0] {
	case clean:
		c.print(fmt.Sprintf("%s\n\n", c.client.CleanCache()))
	case ls:
		for _, item := range c.client.ListCache() {
			c.print(fmt.Sprintf("%s\n", item))
		}
		c.print("\n")
	default:
		c.print(fmt.Sprintf("cache %s: unknown command\n\n", args[0]))
		c.print(fmt.Sprintf("%s\n\n", descriptor))
	}

	return nil
}
