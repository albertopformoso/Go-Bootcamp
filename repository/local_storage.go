package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
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

func (l LocalStorage) Read() ([]model.Pokemon, error) {
    syscall.Umask(0)
	filePath := path.Join(dir, file)
    f, err := os.Open(filePath)
    defer func() {
        if err := f.Close(); err != nil {
            log.Printf("ERROR: file not closed")
        }
    }()
    if err != nil {
        return nil, err
    }

    r := csv.NewReader(f)
    records, err := r.ReadAll()
    if err != nil {
        return nil, err
    }

    pokemons, err := parseCSVData(records)
    if err != nil {
        return nil, err
    }

    return pokemons, nil
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

func parseCSVData(records [][]string) ([]model.Pokemon, error) {
    var pokemons []model.Pokemon
    for i, record := range records {
        if i == 0 {
            continue
        }

        id, err := strconv.Atoi(record[0])
        if err != nil {
            return nil, err
        }

        height, err := strconv.Atoi(record[2])
        if err != nil {
            return nil, err
        }

        weight, err := strconv.Atoi(record[3])
        if err != nil {
            return nil, err
        }

        pokemon := model.Pokemon{
            ID: id,
            Name: record[1],
            Height: height,
            Weight: weight,
            FlatAbilityURLs: record[4],
        }
        pokemons = append(pokemons, pokemon)
    }

    return pokemons, nil
}
