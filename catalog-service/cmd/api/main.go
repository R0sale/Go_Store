package main

import (
	"catalog-service/internal/config"
	"catalog-service/internal/db"
	"catalog-service/internal/handlers"
	"catalog-service/internal/service"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	cfg, err := config.LoadConfig(`C:\go\go_store\catalog-service\config`)
	if err != nil {
		log.Fatal(err.Error())
	}

	database, err := db.ConfigureDb(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	repository := db.NewRepository(database)

	service := service.NewService(repository)

	catalogHandler := handlers.NewCatalogHandler(service)

	mux.HandleFunc("GET /api/catalog/{userId}", catalogHandler.HandleGetCatalog)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Print("Couldnt run the server")
	}
}
