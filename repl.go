package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/foxtrottwist/pokego/fetch"
)

const PROMPT = "Pokédex > "

type config struct {
	next     *string
	previous *string
	client   fetch.Client
}

func start() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands()

	config := &config{
		client: fetch.NewClient(5 * time.Second),
	}

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		cmdName := strings.Fields(line)[0]

		if cmd, exist := cmds[cmdName]; exist {
			err := cmd.run(config)
			if err != nil {
				fmt.Printf("%v\n\n", err)
			}
			continue
		} else {
			fmt.Printf("%s: unknown command\n", cmdName)
			fmt.Print("use `help' for usage.\n\n")
		}
	}
}
