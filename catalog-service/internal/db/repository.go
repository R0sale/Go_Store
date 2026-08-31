package db

import "database/sql"

type repository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) *repository {
	return &repository{
		db: database,
	}
}
