package commands

import "os"

func exitCommand(*config, ...string) error {
	os.Exit(0)
	return nil
}
