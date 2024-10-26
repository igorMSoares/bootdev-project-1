package cmd

import (
	pokeapi "github.com/igorMSoares/bootdev-project-1/internal/pokeapi"
	"github.com/igorMSoares/bootdev-project-1/internal/pokedex"
)

type Config struct {
	PokeapiClient   *pokeapi.Client
	NextLocationUrl *string
	PrevLocationUrl *string
	Pokedex         *pokedex.Pokedex
}

type command struct {
	Name        string
	Description string
	Callback    func(cfg *Config, params ...string) error
}

func GetCommands() map[string]command {
	return map[string]command{
		"map": {
			Name:        "map",
			Description: "Get next location areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "map",
			Description: "Get previous location areas",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Show a list of all the Pokémon in a given area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Try to add a pokemon to the pokedex",
			Callback:    commandCatch,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Print a list of all the names of the Pokemon in the pokedex",
			Callback:    commandPokedex,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Print pokemon details",
			Callback:    commandInspect,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}

}
