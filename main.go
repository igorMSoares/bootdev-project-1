package main

import (
	"time"

	"github.com/igorMSoares/bootdev-project-1/cmd"
	"github.com/igorMSoares/bootdev-project-1/internal/pokeapi"
	"github.com/igorMSoares/bootdev-project-1/internal/pokedex"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &cmd.Config{
		PokeapiClient: pokeapiClient,
		Pokedex:       pokedex.NewPokedex(),
	}

	start(cfg)
}
