package db

import (
	"context"
	"user-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (r repository) AddUser(ctx context.Context, user models.User) (models.User, error) {
	query := `INSERT INTO users (name, password, email, image_url) VALUES($1, $2, $3, $4)
	RETURNING id`

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	password := string(bytes)

	err = r.db.QueryRowContext(ctx, query, user.Name, password, user.Email, user.ImageUrl).Scan(&user.Id)
	if err != nil {
		return models.User{}, err
	}

	return user, err
}
