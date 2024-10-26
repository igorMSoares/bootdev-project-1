package cmd

import (
	"fmt"
)

func commandMap(cfg *Config, params ...string) error {
	res, err := cfg.PokeapiClient.GetLocationAreas(cfg.NextLocationUrl)
	if err != nil {
		return err
	}

	cfg.NextLocationUrl = res.Next
	cfg.PrevLocationUrl = res.Previous

	for _, l := range res.Results {
		fmt.Println(l.Name)
	}

	return nil
}

func commandMapb(cfg *Config, params ...string) error {
	if cfg.PrevLocationUrl == nil {
		return fmt.Errorf("No previous location areas. Use command 'map' instead.")
	}

	res, err := cfg.PokeapiClient.GetLocationAreas(cfg.PrevLocationUrl)
	if err != nil {
		return err
	}

	cfg.NextLocationUrl = res.Next
	cfg.PrevLocationUrl = res.Previous

	for _, l := range res.Results {
		fmt.Println(l.Name)
	}

	return nil
}
