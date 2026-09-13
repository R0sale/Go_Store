package handlers

import (
	"context"
	"user-service/internal/dto"
	"user-service/internal/models"
)

type service interface {
	AddUser(ctx context.Context, user dto.CreateUserDto) error
	LoginUser(ctx context.Context, user dto.LoginUserDto) (models.User, string, error)
}
