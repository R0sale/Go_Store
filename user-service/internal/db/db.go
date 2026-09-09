package db

import (
	"database/sql"
	"fmt"
	"user-service/internal/config"
)

func ConfigureDb(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(`host=%s port=%d password=%s user=%s dbname=%s`, cfg.Database.Host, cfg.Database.Port, cfg.Database.Password, cfg.Database.User, cfg.Database.DbName)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("couldn't open the sql driver")
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("couldn't ping the db, %v", err)
	}

	return db, nil
}
