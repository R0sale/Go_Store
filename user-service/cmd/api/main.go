package main

import (
	"log"
	"net/http"
	"user-service/internal/config"
	"user-service/internal/db"
)

func main() {
	mux := http.NewServeMux()

	cfg, err := config.LoadConfig("C:\\go\\go_store\\user-service\\internal\\config\\config.go")
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := db.ConfigureDb(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
