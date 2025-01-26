package usecase

import (
	"context"

	"next-learn-go/entity"
	"next-learn-go/repository"
)

type CustomerUseCase interface {
	GetAllCustomers() ([]entity.GetAllCustomerResponse, error)
	GetFilteredCustomers(query string) ([]entity.GetFilteredCustomerResponse, error)
	GetCustomerCount() (int, error)
}

type customerUseCase struct {
	cr repository.CustomerRepository
}

func NewCustomerUseCase(cr repository.CustomerRepository) CustomerUseCase {
	return &customerUseCase{cr}
}

func (cu *customerUseCase) GetAllCustomers() ([]entity.GetAllCustomerResponse, error) {
	customers := []entity.Customer{}
	if err := cu.cr.GetAllCustomers(context.Background(), &customers); err != nil {
		return nil, err
	}
	resCustomers := []entity.GetAllCustomerResponse{}
	for _, v := range customers {
		c := entity.GetAllCustomerResponse{}
		c.ID = v.ID
		c.Name = v.Name
		resCustomers = append(resCustomers, c)
	}

	return resCustomers, nil
}

func (cu *customerUseCase) GetFilteredCustomers(query string) ([]entity.GetFilteredCustomerResponse, error) {
	customers := []entity.Customer{}
	if err := cu.cr.GetFilteredCustomers(context.Background(), &customers, query); err != nil {
		return nil, err
	}

	resCustomers := []entity.GetFilteredCustomerResponse{}
	for _, v := range customers {
		c := entity.GetFilteredCustomerResponse{}
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

func (cu *customerUseCase) GetCustomerCount() (int, error) {
	count, err := cu.cr.GetCustomerCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}
