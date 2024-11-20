package commands

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

func catchCommand(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("a pokemon name must be provided")
	}

	pokemon, err := c.client.GetPokemon(args[0])
	if err != nil {
		return err
	}

	c.print(fmt.Sprintf("Throwing a Pokeball at %s...\n", pokemon.Name))
	if tryCatch(pokemon.BaseExperience) {
		c.pokedex[pokemon.Name] = pokemon
		c.print(fmt.Sprintf("%s was caught!\n\n", pokemon.Name))
	} else {
		c.print(fmt.Sprintf("%s escaped!\n\n", pokemon.Name))
	}

	return nil
}

func tryCatch(exp int) bool {
	time.Sleep(1 * time.Second)
	difficulty := 10

	if exp >= 50 {
		difficulty = exp / 5
	}
	if exp >= 150 {
		difficulty = exp / 6
	}

	hp := 0
	for range 11 {
		points := rand.IntN(difficulty)
		hp += points
		if hp > exp {
			return true
		}
	}
	return false
}
