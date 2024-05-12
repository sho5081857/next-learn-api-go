package usecase

import (
	"context"

	"next-learn-go/model"
	"next-learn-go/repository"
)

type ICustomerUsecase interface {
	GetAllCustomers() ([]model.GetAllCustomerResponse, error)
	GetFilteredCustomers(query string) ([]model.GetFilteredCustomerResponse, error)
	GetCustomerCount() (int, error)
}

type customerUsecase struct {
	cr repository.ICustomerRepository
}

func NewCustomerUsecase(cr repository.ICustomerRepository) ICustomerUsecase {
	return &customerUsecase{cr}
}

func (cu *customerUsecase) GetAllCustomers() ([]model.GetAllCustomerResponse, error) {
	customers := []model.Customer{}
	if err := cu.cr.GetAllCustomers(context.Background(), &customers); err != nil {
		return nil, err
	}
	resCustomers := []model.GetAllCustomerResponse{}
	for _, v := range customers {
		c := model.GetAllCustomerResponse{}
		c.ID = v.ID
		c.Name = v.Name
		resCustomers = append(resCustomers, c)
	}

	return resCustomers, nil
}

func (cu *customerUsecase) GetFilteredCustomers(query string) ([]model.GetFilteredCustomerResponse, error) {
	customers := []model.Customer{}
	if err := cu.cr.GetFilteredCustomers(context.Background(), &customers, query); err != nil {
		return nil, err
	}

	resCustomers := []model.GetFilteredCustomerResponse{}
	for _, v := range customers {
		c := model.GetFilteredCustomerResponse{}
		c.ID = v.ID
		c.Name = v.Name
		c.Email = v.Email
		c.ImageUrl = v.ImageUrl
		c.TotalInvoices = v.TotalInvoices
		c.TotalPending = v.TotalPending
		c.TotalPaid = v.TotalPaid
		resCustomers = append(resCustomers, c)
	}

	return resCustomers, nil
}

func (cu *customerUsecase) GetCustomerCount() (int, error) {
	count, err := cu.cr.GetCustomerCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}
