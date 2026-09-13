package services

import (
	"context"
	"time"
	"user-service/internal/dto"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func (s service) LoginUser(ctx context.Context, user dto.LoginUserDto) (models.User, string, error) {
	repoUser := models.User{}

	repoUser.Email = user.Email
	repoUser.Password = user.Password

	dbUser, err := s.repository.LoginUser(ctx, repoUser)
	if err != nil {
		return models.User{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": dbUser.Email,
		"image": dbUser.ImageUrl,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.SecretKey.Key))
	if err != nil {
		return models.User{}, "", err
	}

	return dbUser, tokenString, err
}
