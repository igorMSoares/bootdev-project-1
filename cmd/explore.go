package cmd

import "fmt"

func commandExplore(cfg *Config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Location name not provided")
	}

	fmt.Printf("Exploring %s...\n", params[0])

	location, err := cfg.PokeapiClient.GetLocationByName(params[0])
	if err != nil {
		return err
	}

	if len(location.Encounters) == 0 {
		fmt.Println("No Pokémon found")
		return nil
	}

	fmt.Println("Found Pokémon:")

	for _, encounter := range location.Encounters {
		fmt.Println("  -", encounter.Pokemon.Name)
	}

	return nil
}
