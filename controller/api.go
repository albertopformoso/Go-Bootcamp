package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type API struct {
	fetcher
}

func NewAPI(fetcher fetcher) API {
	return API{fetcher}
}

type fetcher interface {
	Fetch(from, to int) error
}

func (api API) FillCSV(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	requestBody := struct {
		From int `json:"from"`
		To   int `json:"to"`
	}{1, 10}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(
			ErrMessage("Error:", err),
		)
		return
	}

	json.Unmarshal(reqBody, &requestBody)
	if err := api.Fetch(requestBody.From, requestBody.To); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(
			ErrMessage("Error:", err),
		)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		Message("Pokemons Fetched Successfully!"),
	)
}
