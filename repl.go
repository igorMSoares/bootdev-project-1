package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/igorMSoares/bootdev-project-1/cmd"
)

func start(cfg *cmd.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error reading from stdin: %v", err)
			continue
		}

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		cmdName := input[0]

		command, ok := cmd.GetCommands()[cmdName]
		if !ok {
			fmt.Print("Invalid command. Try 'help' for usage info.\n\n")
			continue
		}

		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		err := command.Callback(cfg, args...)
		if err != nil {
			fmt.Printf("> Error running %s command\n%v\n", command.Name, err)
			continue
		}

	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
