package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

type API struct {
	fetcher
	getter
}

func NewAPI(fetcher fetcher, getter getter) API {
	return API{fetcher, getter}
}

type fetcher interface {
	Fetch(from, to int) error
}

type getter interface {
	GetPokemons() ([]model.Pokemon, error)
}

func (api API) FillCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			Message(fmt.Sprintf("Method %v not allowed", r.Method)),
		)
		return
	}

	requestBody := struct {
		From int `json:"from"`
		To   int `json:"to"`
	}{1, 10}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			ErrMessage("Error:", err),
		)
		return
	}

	if err := json.Unmarshal(reqBody, &requestBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(
			ErrMessage("ERROR:", err),
		)
		return
	}

	if err := api.Fetch(requestBody.From, requestBody.To); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(
			ErrMessage("Error:", err),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(
		Message("Pokemons Fetched Successfully!"),
	)
}

func (api API) GetPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			Message(fmt.Sprintf("Method %v not allowed", r.Method)),
		)
		return
	}

	pokemons, err := api.getter.GetPokemons()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(
			ErrMessage("FAIL: can't get the pokemons -", err),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(pokemons)
}
