package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

type LocalStorage struct{}

const (
	file string = "pokemons.csv"
	dir  string = "data"
)

const (
	totalJobs    = 4
	totalWorkers = 30
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
	jobs := make(chan []string, 4)
	results := make(chan model.Pokemon, 4)
	var pokemons []model.Pokemon

	for w := 1; w <= totalWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for _, record := range records {
		jobs <- record
	}
	close(jobs)

	// Receive results
	for a := 1; a <= len(records)-1; a++ {
		result := <-results
		pokemons = append(pokemons, result)
	}
	close(results)

	return pokemons, nil
}

func worker(wId int, jobs <-chan []string, results chan<- model.Pokemon) {
	var wg sync.WaitGroup

	for record := range jobs {
		wg.Add(1)

		go func(record []string) {
			defer wg.Done()
			id, err := strconv.Atoi(record[0])
			if err != nil {
				return
			}

			height, err := strconv.Atoi(record[2])
			if err != nil {
				return
			}

			weight, err := strconv.Atoi(record[3])
			if err != nil {
				return
			}

			pokemon := model.Pokemon{
				ID:              id,
				Name:            record[1],
				Height:          height,
				Weight:          weight,
				FlatAbilityURLs: record[4],
			}
			results <- pokemon
		}(record)
	}

	wg.Wait()
}
