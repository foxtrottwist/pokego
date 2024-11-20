package commands

type command struct {
	name        string
	description string
	Run         func(*config, ...string) error
}

func Commands() map[string]command {
	return map[string]command{
		"cache": {
			name:        "cache <command>",
			description: cacheCmdDesciption,
			Run:         cacheCommand,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Attempts to catch a Pokémon",
			Run:         catchCommand,
		},
		"clear": {
			name:        "clear",
			description: "Clears the PokéGo output",
			Run:         clearCommand,
		},
		// 'exit' command is handled in the repl execution loop to allow for defer term.Cleanup to run.
		"exit": {
			name:        "exit",
			description: "Exits PokéGo",
			Run:         func(c *config, s ...string) error { return nil },
		},
		"explore": {
			name:        "explore <location-area>",
			description: "Displays a list of Pokémon found in a location area",
			Run:         exploreCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Run:         helpCommand,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Displays the attributes of a Pokémon in your Pokédex",
			Run:         inspectCommand,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas",
			Run:         mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas",
			Run:         mapbCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of Pokémon in your Pokédex",
			Run:         pokedexCommand,
		},
	}
}
