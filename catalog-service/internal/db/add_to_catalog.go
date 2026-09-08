package db

import (
	"context"
)

func (r repository) AddToCatalog(ctx context.Context, name string, price float32, imageUrl string) error {
	query := `INSERT INTO catalog (name, price, image_url) VALUES($1, $2, $3)`

	_, err := r.db.ExecContext(ctx, query, &name, &price, &imageUrl)

	return err
}
