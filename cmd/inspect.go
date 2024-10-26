package cmd

import "fmt"

func commandInspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Pokemon name not provided")
	}

	pokemonName := args[0]

	pokemon := cfg.Pokedex.Get(pokemonName)
	if pokemon == nil {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %v: %v\n", stat.Stat.Name, stat.Base)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %v\n", t.Type.Name)
	}

	return nil
}
