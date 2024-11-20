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

	for key, value := range cmds {
		if l := len(value.name); l > padding {
			padding = l
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		cmd, _ := cmds[key]
		c.print(fmt.Sprintf("%s:%s%s\n", cmd.name, strings.Repeat(" ", padding+2-len(cmd.name)), cmd.description))
	}

	c.print("\n")
	return nil
}
