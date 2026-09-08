package service

import (
	"catalog-service/internal/models"
	"context"
)

type repo interface {
	GetCatalog(ctx context.Context) ([]models.Item, error)
	AddToCatalog(ctx context.Context, name string, price float32, imageUrl string) error
}
