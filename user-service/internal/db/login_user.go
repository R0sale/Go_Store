package db

import (
	"context"
	"time"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (r repository) LoginUser(ctx context.Context, user models.User) (string, error) {
	var exist bool
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 AND password = $2)`

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	if err := r.db.QueryRowContext(ctx, query, user.Email, passwordHash).Scan(&exist); err != nil {
		return "", err
	}

	getUserQuery := `SELECT image_url FROM users WHERE email = $1`
	if err := r.db.QueryRowContext(ctx, getUserQuery, user.Email).Scan(&user.ImageUrl); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"image": user.ImageUrl,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(r.cfg.SecretKey.Key))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
