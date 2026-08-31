package handlers

import (
	"catalog-service/internal/models"
	"context"
)

type Service interface {
	GetCatalog(ctx context.Context, userId int) ([]models.Item, error)
}
