package commands

type command struct {
	name        string
	description string
	Run         func(*config, ...string) error
}

func Commands() map[string]command {
	return map[string]command{
		"cache": {
			name:        "cache",
			description: "Manipulates the PokéGo cache",
			Run:         cacheCommand,
		},
		"clear": {
			name:        "clear",
			description: "Clears the PokéGo output",
			Run:         clearCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exits PokéGo",
			Run:         exitCommand,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of Pokémon found in a location area",
			Run:         exploreCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Run:         helpCommand,
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
	}
}
