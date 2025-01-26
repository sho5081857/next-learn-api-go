package usecase

import (
	"context"
	"next-learn-go/entity"
	"next-learn-go/repository"
)

type RevenueUseCase interface {
	GetAllRevenues() ([]entity.Revenue, error)
}

type revenueUseCase struct {
	rr repository.RevenueRepository
}

func NewRevenueUseCase(rr repository.RevenueRepository) RevenueUseCase {
	return &revenueUseCase{rr}
}

func (ru *revenueUseCase) GetAllRevenues() ([]entity.Revenue, error) {
	revenues := []entity.Revenue{}
	if err := ru.rr.GetAllRevenues(context.Background(), &revenues); err != nil {
		return nil, err
	}
	return revenues, nil
}
