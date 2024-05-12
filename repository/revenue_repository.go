package repository

import (
	"context"
	"next-learn-go/model"

	"github.com/uptrace/bun"
)


type IRevenueRepository interface {
	GetAllRevenues(ctx context.Context, revenues *[]model.Revenue) error
}

type revenueRepository struct {
	db *bun.DB
}


func NewRevenueRepository(db *bun.DB) IRevenueRepository {
	return &revenueRepository{db}
}

func (rr *revenueRepository) GetAllRevenues(ctx context.Context, revenues *[]model.Revenue) error {
	if err := rr.db.NewSelect().
		Model(revenues).
		Scan(ctx); err != nil {
		return err
	}
	return nil
}

