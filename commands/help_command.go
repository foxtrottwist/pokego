package commands

import (
	"fmt"
	"sort"
	"strings"
)

func helpCommand(c *config, args ...string) error {
	c.print("\nWelcome to the PokéGo Pokédex!\nUsage:\n\n")

	cmds := Commands()
	keys := make([]string, 0, len(cmds))
	padding := 0

	for key := range cmds {
		if l := len(key); l > padding {
			padding = len(key)
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		cmd, _ := cmds[key]
		c.print(fmt.Sprintf("%s:%s%s\n", cmd.name, strings.Repeat(" ", padding*2-len(key)), cmd.description))
	}

	c.print("\n")
	return nil
}
