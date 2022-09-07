package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

type LocalStorage struct{}

const filePath = "data/pokemons.csv"

// Creating a file and writing the data to it.
func (l LocalStorage) Write(pokemons []model.Pokemon) error {
	file, err := os.Create(filePath)
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("ERROR: file not closed")
		}
	}()
	if err != nil {
		return err
	}

	w := csv.NewWriter(file)
	records := buildRecords(pokemons)
	if err := w.WriteAll(records); err != nil {
		return err
	}

	return nil
}

// It takes a slice of Pokemon structs and returns a slice of slices of strings
func buildRecords(pokemons []model.Pokemon) [][]string {
	headers := []string{"id", "name", "height", "weight", "flat_abilities"}
	records := [][]string{headers}
	for _, pokemon := range pokemons {
		record := fmt.Sprintf("%d,%s,%d,%d,%s",
			pokemon.ID,
			pokemon.Name,
			pokemon.Height,
			pokemon.Weight,
			pokemon.FlatAbilitiyURLs,
		)
		records = append(records, strings.Split(record, ","))
	}

	return records
}
