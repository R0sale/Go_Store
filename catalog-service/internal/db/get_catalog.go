package db

import (
	"catalog-service/internal/models"
	"context"
)

func (r *repository) GetCatalog(ctx context.Context) ([]models.Item, error) {
	query := `SELECT id, name, price, image_url FROM catalog`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.Id, &item.Name, &item.Price, &item.ImageUrl)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
