package repository

import (
	"context"
	"next-learn-go/entity"

	"github.com/uptrace/bun"
)

type RevenueRepository interface {
	GetAllRevenues(ctx context.Context, revenues *[]entity.Revenue) error
}

type revenueRepository struct {
	db *bun.DB
}

func NewRevenueRepository(db *bun.DB) RevenueRepository {
	return &revenueRepository{db}
}

func (rr *revenueRepository) GetAllRevenues(ctx context.Context, revenues *[]entity.Revenue) error {
	if err := rr.db.NewSelect().
		Model(revenues).
		Scan(ctx); err != nil {
		return err
	}
	return nil
}
