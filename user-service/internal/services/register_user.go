package services

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/dto"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func (s service) AddUser(ctx context.Context, user dto.CreateUserDto) (models.User, string, error) {
	repoUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
		Password: user.Password,
	}

	newUser, err := s.repository.AddUser(ctx, repoUser)
	if err != nil {
		return models.User{}, "", err
	}

	s.wg.Add(1)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("couldnt send the email err: %s\n", err)
			}
		}()

		defer s.wg.Done()

		err := s.mailer.Send(user.Email, "user_welcome.tmpl", user)
		if err != nil {
			fmt.Printf("couldnt send the email err: %s\n", err.Error())
		}
	}()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": newUser.Email,
		"image": newUser.ImageUrl,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.SecretKey.Key))
	if err != nil {
		return models.User{}, "", err
	}

	return newUser, tokenString, nil
}
