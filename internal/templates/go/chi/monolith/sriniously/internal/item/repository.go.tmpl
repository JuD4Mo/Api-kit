package item

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) Repository {
	return &repo{pool: pool}
}

func (r *repo) Create(ctx context.Context, item *Item) error {
	query := `
		INSERT INTO items (id, name, description, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.pool.Exec(ctx, query,
		item.ID,
		item.Name,
		item.Description,
		item.Status,
		item.CreatedAt,
	)
	return err
}

func (r *repo) GetByID(ctx context.Context, id string) (*Item, error) {
	query := `
		SELECT id, name, description, status, created_at 
		FROM items 
		WHERE id = $1
	`
	var item Item
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&item.ID,
		&item.Name,
		&item.Description,
		&item.Status,
		&item.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &item, nil
}
