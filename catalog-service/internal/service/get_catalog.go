package service

import (
	"catalog-service/internal/models"
	"context"
)

func (s *service) GetCatalog(ctx context.Context, userId int) ([]models.Item, error) {
	items, err := s.repository.GetCatalog(ctx, userId)
	return items, err
}
