package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/foxtrottwist/pokego/commands"
	"github.com/foxtrottwist/pokego/terminal"
)

const PROMPT = "Pokédex > "

func start() {
	term := terminal.New(PROMPT)
	defer term.Cleanup()

	cmds := commands.Commands()
	config := commands.NewConfig(term.Write, 5*time.Second, 5*time.Minute)

	for {
		line, err := term.ReadLine()
		if err == io.EOF {
			break
		}

		fields := strings.Fields(strings.ToLower(line))
		if len(fields) == 0 {
			continue
		}

		cmdName := fields[0]
		if cmdName == "exit" {
			break
		}

		if cmd, exist := cmds[cmdName]; exist {
			args := fields[1:]
			err := cmd.Run(config, args...)
			if err != nil {
				term.Write(fmt.Sprintf("%v\n\n", err))
			}
			continue
		} else {
			term.Write(fmt.Sprintf("%s: unknown command\n", cmdName))
			term.Write("use `help' for usage.\n\n")
		}
	}
}
