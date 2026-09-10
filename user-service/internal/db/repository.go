package db

import (
	"database/sql"
	"user-service/internal/config"
)

type repository struct {
	db  *sql.DB
	cfg config.Config
}

func NewRepository(database *sql.DB, config config.Config) *repository {
	return &repository{
		db:  database,
		cfg: config,
	}
}
