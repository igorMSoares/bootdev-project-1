package pokedex

import "github.com/igorMSoares/bootdev-project-1/internal/pokeapi"

type Pokedex map[string]pokeapi.Pokemon

func NewPokedex() *Pokedex {
	return &Pokedex{}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	(*p)[pokemon.Name] = pokemon
}

func (p *Pokedex) Get(name string) *pokeapi.Pokemon {
	if p, ok := (*p)[name]; ok {
		return &p
	}

	return nil
}
