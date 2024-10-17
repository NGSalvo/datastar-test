package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ngsalvo/datastar-test/handlers"
)

func main() {
	router := chi.NewRouter()

	handlers.SetupHome(router)

	log.Printf("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
