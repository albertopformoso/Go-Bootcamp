package service

import (
	"context"
	"log"
	"sort"
	"strings"
	"sync"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

var wg = sync.WaitGroup{}

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

type PokemonsByID []model.Pokemon

func (x PokemonsByID) Len() int           { return len(x) }
func (x PokemonsByID) Less(i, j int) bool { return x[i].ID < x[j].ID }
func (x PokemonsByID) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Result struct {
	Error   error
	Pokemon model.Pokemon
}

func NewFetcher(api api, storage writer) Fetcher {
	return Fetcher{api, storage}
}

// Fetching a range of pokemon from the API, and writing them to the storage.
func (f Fetcher) Fetch(from, to int) error {
	var pokemons []model.Pokemon
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    channel := GetResponses(ctx, from, to, f)

	for id := from; id <= to; id++ {
		result := <-channel
        if result.Error != nil {
            log.Printf("ERROR: %v", result.Error)
            cancel()
            continue
        }
		pokemons = append(pokemons, result.Pokemon)
	}

    sort.Sort(PokemonsByID(pokemons))

	return f.storage.Write(pokemons)
}

// It takes a context, a range of IDs, and a Fetcher, and returns a channel of Results
func GetResponses(ctx context.Context, from, to int, f Fetcher) <-chan Result {
    results := make(chan Result)

    go func() {
        wg.Wait()
        close(results)
    }()

    for id := from; id <= to; id++ {
        wg.Add(1)
        go PingAPI(ctx, id, f, results)
    }

    return results
}

// It takes a context, a waitgroup, an id, a fetcher, and a channel of results. It then fetches a
// pokemon from the API, flattens the abilities, and sends the result to the channel
func PingAPI(ctx context.Context, id int, f Fetcher, results chan Result) {
	defer wg.Done()
	var result Result
	pokemon, err := f.api.FetchPokemon(id)

	var flatAbilities []string
	for _, t := range pokemon.Abilities {
		flatAbilities = append(flatAbilities, t.Ability.URL)
	}
	pokemon.FlatAbilityURLs = strings.Join(flatAbilities, "|")

	result = Result{Error: err, Pokemon: pokemon}

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	case results <- result:
	}
}

