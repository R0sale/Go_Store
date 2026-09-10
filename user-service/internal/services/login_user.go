package services

import (
	"context"
	"user-service/internal/dto"
	"user-service/internal/models"
)

func (s service) LoginUser(ctx context.Context, user dto.LoginUserDto) (string, error) {
	repoUser := models.User{}

	repoUser.Email = user.Email
	repoUser.Password = user.Password

	token, err := s.repository.LoginUser(ctx, repoUser)
	return token, err
}
