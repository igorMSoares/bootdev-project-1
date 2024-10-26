package cmd

import "fmt"

func commandPokedex(cfg *Config, args ...string) error {
	if len((*cfg.Pokedex)) == 0 {
		fmt.Println("Pokedex is empty.")
		fmt.Println("You haven't caught any Pokemon yet.")
	}

	fmt.Println("Your Pokedex:")
	for name := range *cfg.Pokedex {
		fmt.Printf("  - %s\n", name)
	}

	return nil
}
