package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

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
	GetEvenOdd(ty string, items, itemsPerWorkers int) ([]model.Pokemon, error)
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

func (api API) GetEvenOdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(
			Message(fmt.Sprintf("Method %v not allowed", r.Method)),
		)
		return
	}

	typeKey, ok := r.URL.Query()["type"]
    if !ok || len(typeKey[0]) < 1 {
        w.WriteHeader(http.StatusBadRequest)
        log.Println("URL Param 'type' is missing")
        _ = json.NewEncoder(w).Encode(
            Message("URL Param 'type' is missing"),
        )
        return
    }

    itemsKey, ok := r.URL.Query()["items"]
    if !ok || len(itemsKey[0]) < 1 {
        w.WriteHeader(http.StatusBadRequest)
        log.Println("URL Param 'items' is missing")
        _ = json.NewEncoder(w).Encode(
            Message("URL Param 'items' is missing"),
        )
        return
    }

    itemsPerWorkersKey, ok := r.URL.Query()["items_per_workers"]
    if !ok || len(itemsPerWorkersKey[0]) < 1 {
        w.WriteHeader(http.StatusBadRequest)
        log.Println("URL Param 'items_per_workers' is missing")
        _ = json.NewEncoder(w).Encode(
            Message("URL Param 'items_per_workers' is missing"),
        )
        return
    }

    ty := typeKey[0]
    items, err := strconv.Atoi(itemsKey[0])
    if err != nil {
        _ = json.NewEncoder(w).Encode(
			ErrMessage("ERROR:", err),
		)
        return
    }
    itemsPerWorkers, err := strconv.Atoi(itemsPerWorkersKey[0])
    if err != nil {
        _ = json.NewEncoder(w).Encode(
			ErrMessage("ERROR:", err),
		)
        return
    }

	pokemons, err := api.getter.GetEvenOdd(ty, items, itemsPerWorkers)
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
