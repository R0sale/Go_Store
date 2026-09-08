package handlers

import (
	"catalog-service/internal/dto"
	"catalog-service/internal/models"
	"context"
)

type Service interface {
	GetCatalog(ctx context.Context) ([]models.Item, error)
	AddToCatalog(ctx context.Context, item dto.CreateItemDto) error
}
