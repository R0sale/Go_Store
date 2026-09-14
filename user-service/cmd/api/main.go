package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"user-service/internal/config"
	"user-service/internal/db"
	"user-service/internal/handlers"
	"user-service/internal/mailer"
	"user-service/internal/services"

	"github.com/rs/cors"
)

var wg sync.WaitGroup

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

	mailer := mailer.New(cfg.Smtp.Host, cfg.Smtp.Port, cfg.Smtp.Username, cfg.Smtp.Password, "kvusov@bk.ru")

	repository := db.NewRepository(database)
	service := services.NewUserService(repository, mailer, *cfg, &wg)
	handler := handlers.NewUserHandler(service)

	mux.HandleFunc("POST /api/users/register", handler.HandleAddUser)
	mux.HandleFunc("POST /api/users/login", handler.HandleLoginUser)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	serverHandler := c.Handler(mux)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      serverHandler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit
		fmt.Println(s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdown <- server.Shutdown(ctx)
	}()

	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}

	wg.Wait()
	err = <-shutdown
	if err != nil {
		log.Fatal(err.Error())
	}
}
