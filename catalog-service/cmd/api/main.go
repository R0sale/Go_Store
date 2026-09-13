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
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/rs/cors"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func main() {
	mux := http.NewServeMux()

	cfg, err := config.LoadConfig()
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

	mux.HandleFunc("GET /api/catalog/", catalogHandler.HandleGetCatalog)
	mux.HandleFunc("POST /api/catalog", catalogHandler.HandleAddToCatalog)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Print("Couldnt run the server")
	}
}
