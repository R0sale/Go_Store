package main

import (
	"catalog-service/internal/config"
	"catalog-service/internal/db"
	"catalog-service/internal/handlers"
	"catalog-service/internal/service"
	"catalog-service/internal/validators"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

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

	itemValidator := validators.NewItemValidator(validate)
	catalogHandler := handlers.NewCatalogHandler(service, itemValidator)

	mux.HandleFunc("GET /api/catalog/{userId}", catalogHandler.HandleGetCatalog)
	mux.HandleFunc("POST /api/catalog", catalogHandler.HandleAddToCatalog)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Print("Couldnt run the server")
	}
}
