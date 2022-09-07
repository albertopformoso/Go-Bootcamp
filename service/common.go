package service

import "github.com/albertopformoso/Go-Bootcamp/model"

type api interface {
	FetchPokemon(id int) (model.Pokemon, error)
	GetPokemons() ([]model.Pokemon, error)
	GetEvenOdd(ty string, items, itemsPerWorkers int) ([]model.Pokemon, error)
}
