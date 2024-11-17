package commands

import (
	"os"
	"os/exec"
	"runtime"
)

func clearCommand(*config, ...string) error {
	if runtime.GOOS == "windows" {
		return clearHelper(exec.Command("cmd", "/c", "cls"))
	}
	return clearHelper(exec.Command("clear"))
}

func clearHelper(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
