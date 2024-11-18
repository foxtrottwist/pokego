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
	if len(args) == 0 {
		return errors.New(
			fmt.Sprintf("no `cache` command provided\nManipulates the PokéGo cache\n\nUsage: \n\ncache %s: removes all items from the cache\ncache %s: lists all items in the cache", clean, ls),
		)
	}

	switch args[0] {
	case clean:
		c.write(fmt.Sprintf("%s\n", c.client.CleanCache()))
		c.write("\n")
	case ls:
		for _, item := range c.client.ListCache() {
			c.write(fmt.Sprintf("%s\n", item))
		}
		c.write("\n")
	default:
		c.write(fmt.Sprintf("cache %s: unknown command\n", args[0]))
		c.write("\n")
	}

	return nil
}
