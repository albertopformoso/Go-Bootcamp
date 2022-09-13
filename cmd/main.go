package main

import (
	"log"
	"net/http"
	"os"

	"github.com/albertopformoso/Go-Bootcamp/controller"
	"github.com/albertopformoso/Go-Bootcamp/repository"
	"github.com/albertopformoso/Go-Bootcamp/router"
	"github.com/albertopformoso/Go-Bootcamp/service"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	api := repository.PokeAPI{}
	storage := repository.LocalStorage{}

	svcFetcher := service.NewFetcher(api, storage)
	svcGetter := service.NewGetter(api)
	ctrl := controller.NewAPI(svcFetcher, svcGetter)

	router.Routes(mux, ctrl)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
