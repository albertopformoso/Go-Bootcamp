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

	svc := service.NewFetcher(api, storage)
	ctrl := controller.NewAPI(svc)

	router.Routes(mux, ctrl)

	log.Println("Listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
