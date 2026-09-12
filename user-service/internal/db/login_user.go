package db

import (
	"context"
	"time"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (r repository) LoginUser(ctx context.Context, user models.User) (string, error) {
	var passwordHash string
	query := `SELECT password FROM users WHERE email = $1`

	if err := r.db.QueryRowContext(ctx, query, user.Email).Scan(&passwordHash); err != nil {
		return "", err
	}

	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(user.Password))
	if err != nil {
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
