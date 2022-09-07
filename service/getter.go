package service

import (
	"sort"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

type Getter struct {
	api api
}

func NewGetter(api api) Getter {
	return Getter{api}
}

func (g Getter) GetPokemons() ([]model.Pokemon, error) {
	pokemons, err := g.api.GetPokemons()
	if err != nil {
		return nil, err
	}

	sort.Sort(PokemonsByID(pokemons))
	return pokemons, nil
}

func (g Getter) GetEvenOdd(ty string, items, itemsPerWorkers int) ([]model.Pokemon, error) {
	pokemons, err := g.api.GetEvenOdd(ty, items, itemsPerWorkers)
	if err != nil {
		return nil, err
	}

	sort.Sort(PokemonsByID(pokemons))
	return pokemons, nil
}
