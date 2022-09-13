package repository

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"

	"github.com/albertopformoso/Go-Bootcamp/model"
)

func TestRead(t *testing.T) {
	storage := LocalStorage{}

	table := []struct {
		name   string
		values []model.Pokemon
		err    error
	}{
		{name: "Error test", values: nil, err: &fs.PathError{
			Op:   "open",
			Path: "data/pokemons.csv",
			Err:  errors.New("no such file or directory"),
		}},
	}

	for _, v := range table {
		t.Run(v.name, func(t *testing.T) {
			pokemons, gotErr := storage.Read()

			if reflect.TypeOf(gotErr) != reflect.TypeOf(v.err) {
				t.Errorf(
					"Error: expected %q, obtained %q",
					reflect.TypeOf(gotErr),
					reflect.TypeOf(v.err),
				)
			}

			if gotErr.Error() != v.err.Error() {
				t.Errorf("Error: expected %q, obtained %q", gotErr.Error(), v.err.Error())
			}

			if len(pokemons) != len(v.values) {
				t.Errorf("Error: expected %q, obtained %q", pokemons, v.values)
			}
		})
	}
}
