package cmd

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("Pokemon name not provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.PokeapiClient.GetPokemonByName(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a ball at %s...\n", pokemonName)

	baseVal := 40
	randomVal := rand.IntN(pokemon.Experience)
	if randomVal > baseVal {
		fmt.Printf("%s escaped!\n", pokemonName)
		return nil
	}

	cfg.Pokedex.Add(*pokemon)
	fmt.Printf("%s was caught!\n", pokemonName)

	return nil
}
