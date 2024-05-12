package usecase

import (
	"context"
	"next-learn-go/model"
	"next-learn-go/repository"
)

type IRevenueUsecase interface {
	GetAllRevenues() ([]model.Revenue, error)
}

type revenueUsecase struct {
	rr repository.IRevenueRepository
}

func NewRevenueUsecase(rr repository.IRevenueRepository) IRevenueUsecase {
	return &revenueUsecase{rr}
}

func (ru *revenueUsecase) GetAllRevenues() ([]model.Revenue, error) {
	revenues := []model.Revenue{}
	if err := ru.rr.GetAllRevenues(context.Background(), &revenues); err != nil {
		return nil, err
	}
	return revenues, nil
}
