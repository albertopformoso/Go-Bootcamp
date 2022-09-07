package service

import (
	"strings"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

type api interface {
	FetchPokemon(id int) (model.Pokemon, error)
}

type writer interface {
	Write(pokemons []model.Pokemon) error
}

type Fetcher struct {
	api     api
	storage writer
}

func NewFetcher(api api, storage writer) Fetcher {
	return Fetcher{api, storage}
}

func (f Fetcher) Fetch(from, to int) error {
	var pokemons []model.Pokemon
	for id := from; id <= to; id++ {
		pokemon, err := f.api.FetchPokemon(id)
		if err != nil {
			return err
		}

		var flatAbilities []string
		for _, t := range pokemon.Abilities {
			flatAbilities = append(flatAbilities, t.Ability.URL)
		}
		pokemon.FlatAbilitiyURLs = strings.Join(flatAbilities, "|")

		pokemons = append(pokemons, pokemon)
	}

	return f.storage.Write(pokemons)
}
