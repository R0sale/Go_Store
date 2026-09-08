package service

import (
	"catalog-service/internal/dto"
	"context"
)

func (s service) AddToCatalog(ctx context.Context, item dto.CreateItemDto) error {
	return s.repository.AddToCatalog(ctx, item.Name, item.Price, item.ImageUrl)
}
