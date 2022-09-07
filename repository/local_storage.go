package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"syscall"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

type LocalStorage struct{}

const (
	file string = "pokemons.csv"
	dir  string = "data"
)

// Creating a file and writing the data to it.
func (l LocalStorage) Write(pokemons []model.Pokemon) error {
	syscall.Umask(0)
	filePath := path.Join(dir, file)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0777); err != nil {
			log.Println(err)
		}
	}

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
			pokemon.FlatAbilityURLs,
		)
		records = append(records, strings.Split(record, ","))
	}

	return records
}
