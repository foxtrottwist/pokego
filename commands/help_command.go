package commands

import "fmt"

func helpCommand(*config, ...string) error {
	fmt.Printf("\nWelcome to the PokéGo Pokédex!\nUsage:\n\n")

	for _, cmd := range Commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()
	return nil
}
