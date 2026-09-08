package service

import (
	"catalog-service/internal/models"
	"context"
)

func (s *service) GetCatalog(ctx context.Context) ([]models.Item, error) {
	items, err := s.repository.GetCatalog(ctx)
	return items, err
}
