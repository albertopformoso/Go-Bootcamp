package router

import (
	"net/http"

	"github.com/albertopformoso/Go-Bootcamp/controller"
)

func Routes(router *http.ServeMux, ctrl controller.API) {
	router.HandleFunc("/api/v1/provide", ctrl.FillCSV)
    router.HandleFunc("/api/v1/pokemons", ctrl.GetPokemons)
}
