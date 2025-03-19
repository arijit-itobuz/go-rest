package main

import (
	"log"
	"net/http"

	"github.com/arijit-itobuz/go-rest/src/config"

	"github.com/go-chi/chi/v5"
)

func main() {

	log.Printf("[main]: AppConfig: %+v\n", config.AppConfig)

	router := chi.NewRouter()

	server := &http.Server{
		Addr:    ":" + config.AppConfig.PORT,
		Handler: router,
	}

	log.Printf("[main]: Server starting on http://localhost:%v\n", config.AppConfig.PORT)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("[main]: Error %v", err)
	}

}
