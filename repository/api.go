package repository

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

type PokeAPI struct{}

const url = "https://pokeapi.co/api/v2/"

func (pa PokeAPI) FetchPokemon(id int) (model.Pokemon, error) {
	pokemon := model.Pokemon{}

	pokeUrl := fmt.Sprintf("%s/pokemon/%d", url, id)
	resp, err := http.Get(pokeUrl)
	if err != nil {
		log.Println("ERROR:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR:", err)
	}

	// Unmarshal result
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
	}

	return pokemon, nil
}

func (pa PokeAPI) GetPokemons() ([]model.Pokemon, error) {
	ls := LocalStorage{}
	pokemons, err := ls.Read()
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}
