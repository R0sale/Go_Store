package main

import (
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/db"
	"user-service/internal/handlers"
	"user-service/internal/services"
)

func main() {
	mux := http.NewServeMux()

	cfg, err := config.LoadConfig(`C:\go\go_store\user-service\config`)
	if err != nil {
		log.Fatal(err.Error())
	}

	database, err := db.ConfigureDb(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	repository := db.NewRepository(database)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)

	mux.HandleFunc("POST /api/users", handler.HandleAddUser)

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
