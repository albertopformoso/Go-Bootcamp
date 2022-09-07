package service

import "github.com/albertopformoso/Go-Bootcamp/model"

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

	return pokemons, nil
}
