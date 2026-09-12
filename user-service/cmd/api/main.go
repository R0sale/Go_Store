package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"user-service/internal/config"
	"user-service/internal/db"
	"user-service/internal/handlers"
	"user-service/internal/services"

	"github.com/rs/cors"
)

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

	repository := db.NewRepository(database, *cfg)
	service := services.NewUserService(repository)
	handler := handlers.NewUserHandler(service)

	mux.HandleFunc("POST /api/users", handler.HandleAddUser)
	mux.HandleFunc("POST /api/users/login", handler.HandleLoginUser)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	serverHandler := c.Handler(mux)

	fmt.Println(cfg.Server.Port)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      serverHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
