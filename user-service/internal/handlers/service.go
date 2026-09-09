package handlers

import (
	"context"
	"user-service/internal/dto"
)

type service interface {
	AddUser(ctx context.Context, user dto.CreateUserDto) error
}
