package services

import (
	"context"
	"user-service/internal/dto"
	"user-service/internal/models"
)

func (s service) AddUser(ctx context.Context, user dto.CreateUserDto) error {
	repoUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		ImageUrl: user.ImageUrl,
		Password: user.Password,
	}

	err := s.repository.AddUser(ctx, repoUser)

	return err
}
