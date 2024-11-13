package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const PROMPT = "Pokédex > "

func start() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := commands()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		cmdName := strings.Fields(line)[0]

		if cmd, exist := cmds[cmdName]; exist {
			err := cmd.run()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Printf("%s: unknown command\n", cmdName)
			fmt.Print("use `help' for usage.\n\n")
		}
	}
}
