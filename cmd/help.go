package cmd

import "fmt"

func commandHelp(cfg *Config, params ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for cmd, info := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd, info.Description)
	}

	fmt.Println()

	return nil
}
