package service

import (
	"catalog-service/internal/models"
	"context"
)

type repo interface {
	GetCatalog(ctx context.Context) ([]models.Item, error)
}
