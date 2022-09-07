package controller

import (
	"encoding/json"
	"io"
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
