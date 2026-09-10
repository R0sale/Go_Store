package db

import (
	"context"
	"user-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (r repository) AddUser(ctx context.Context, user models.User) error {
	query := `INSERT INTO users (name, password, email, image_url) VALUES($1, $2, $3, $4)`

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	password := string(bytes)

	_, err = r.db.ExecContext(ctx, query, user.Name, password, user.Email, user.ImageUrl)

	return err
}
