package item

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Create(ctx context.Context, name, description string) (*Item, error) {
	item := &Item{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Status:      "active",
		CreatedAt:   time.Now(),
	}

	if err := s.repo.Create(ctx, item); err != nil {
		return nil, fmt.Errorf("creating item: %w", err)
	}

	return item, nil
}

func (s *service) GetByID(ctx context.Context, id string) (*Item, error) {
	item, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting item: %w", err)
	}

	return item, nil
}
