package db

import (
	"context"
	"fmt"
	"user-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (r repository) LoginUser(ctx context.Context, user models.User) (models.User, error) {
	var passwordHash string
	query := `SELECT password FROM users WHERE email = $1`

	if err := r.db.QueryRowContext(ctx, query, user.Email).Scan(&passwordHash); err != nil {
		return models.User{}, err
	}

	fmt.Println(user.Password)

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
	if err != nil {
		return models.User{}, err
	}

	getUserQuery := `SELECT id, name, image_url FROM users WHERE email = $1`
	if err := r.db.QueryRowContext(ctx, getUserQuery, user.Email).Scan(&user.Id, &user.Name, &user.ImageUrl); err != nil {
		return models.User{}, err
	}

	return user, nil
}
