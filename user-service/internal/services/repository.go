package services

import (
	"context"
	"user-service/internal/models"
)

type repository interface {
	AddUser(ctx context.Context, user models.User) error
	LoginUser(ctx context.Context, user models.User) (string, error)
}
