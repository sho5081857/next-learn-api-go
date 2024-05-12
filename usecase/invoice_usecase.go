package usecase

import (
	"context"
	"fmt"
	"next-learn-go/model"
	"next-learn-go/repository"
	"next-learn-go/validator"

	"github.com/google/uuid"
)

type IInvoiceUsecase interface {
	GetLatestInvoices(offset, limit int) ([]model.GetLatestInvoicesResponse, error)
	GetFilteredInvoices(query string, offset, limit int) ([]model.GetFilteredInvoicesResponse, error)
	GetInvoiceCount() (int, error)
	GetInvoiceStatusCount() (int, int, error)
	GetInvoicesPages(query string, offset, limit int) (int, error)
	GetInvoiceById(invoiceId uuid.UUID) (model.GetInvoiceByIdResponse, error)
	CreateInvoice(invoice model.Invoice) (model.InvoiceResponse, error)
	UpdateInvoice(invoice model.Invoice, invoiceId uuid.UUID) (model.InvoiceResponse, error)
	DeleteInvoice(invoiceId uuid.UUID) error
}

type invoiceUsecase struct {
	ir repository.IInvoiceRepository
	iv validator.IInvoiceValidator
}

func NewInvoiceUsecase(ir repository.IInvoiceRepository, iv validator.IInvoiceValidator) IInvoiceUsecase {
	return &invoiceUsecase{ir, iv}
}

func (iu *invoiceUsecase) GetLatestInvoices(offset, limit int) ([]model.GetLatestInvoicesResponse, error) {
	invoices := []model.Invoice{}
	if err := iu.ir.GetLatestInvoices(context.Background(), &invoices, offset, limit); err != nil {
		return nil, err
	}
	resInvoices := []model.GetLatestInvoicesResponse{}
	for _, v := range invoices {
		i := model.GetLatestInvoicesResponse{}
		i.ID = v.ID
		i.Name = v.Customer.Name
		i.ImageUrl = v.Customer.ImageUrl
		i.Email = v.Customer.Email
		i.Amount = v.Amount
		resInvoices = append(resInvoices, i)
	}
	return resInvoices, nil
}

func (iu *invoiceUsecase) GetFilteredInvoices(query string, offset, limit int) ([]model.GetFilteredInvoicesResponse, error) {
	invoices := []model.Invoice{}
	if err := iu.ir.GetFilteredInvoices(context.Background(), &invoices, query, offset, limit); err != nil {
		return nil, err
	}
	resInvoices := []model.GetFilteredInvoicesResponse{}
	for _, v := range invoices {
		i := model.GetFilteredInvoicesResponse{}
		i.ID = v.ID
		i.CustomerId = v.Customer.ID
		i.Name = v.Customer.Name
		i.Email = v.Customer.Email
		i.ImageUrl = v.Customer.ImageUrl
		i.Amount = v.Amount
		i.Date = v.Date
		i.Status = v.Status
		resInvoices = append(resInvoices, i)
	}
	return resInvoices, nil
}

func (iu *invoiceUsecase) GetInvoiceCount() (int, error) {
	count, err := iu.ir.GetInvoiceCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iu *invoiceUsecase) GetInvoiceStatusCount() (int, int, error) {
	pending, paid, err := iu.ir.GetInvoiceStatusCount(context.Background())
	if err != nil {
		return 0, 0, err
	}
	return pending, paid, nil
}

func (iu *invoiceUsecase) GetInvoicesPages(query string, offset, limit int) (int, error) {
	count, err := iu.ir.GetInvoicesPages(context.Background(), query, offset, limit)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iu *invoiceUsecase) GetInvoiceById(invoiceId uuid.UUID) (model.GetInvoiceByIdResponse, error) {
	invoice := model.Invoice{}
	if err := iu.ir.GetInvoiceById(context.Background(), &invoice, invoiceId); err != nil {
		return model.GetInvoiceByIdResponse{}, err
	}

	resInvoice := model.GetInvoiceByIdResponse{}
	resInvoice.ID = invoice.ID
	resInvoice.CustomerId = invoice.Customer.ID
	resInvoice.Amount = invoice.Amount
	resInvoice.Status = invoice.Status

	return resInvoice, nil
}

func (iu *invoiceUsecase) CreateInvoice(invoice model.Invoice) (model.InvoiceResponse, error) {
	if err := iu.iv.InvoiceValidate(invoice); err != nil {
		return model.InvoiceResponse{}, err
	}
	if err := iu.ir.CreateInvoice(context.Background(), &invoice); err != nil {
		return model.InvoiceResponse{}, err
	}

	resInvoice := model.InvoiceResponse{}
	resInvoice.ID = invoice.ID
	resInvoice.Amount = invoice.Amount
	resInvoice.Date = invoice.Date
	resInvoice.Status = invoice.Status
	resInvoice.Customer.Name = invoice.Customer.Name
	resInvoice.Customer.Email = invoice.Customer.Email
	resInvoice.Customer.ImageUrl = invoice.Customer.ImageUrl

	return resInvoice, nil
}

func (iu *invoiceUsecase) UpdateInvoice(invoice model.Invoice, invoiceId uuid.UUID) (model.InvoiceResponse, error) {
	if err := iu.iv.InvoiceValidate(invoice); err != nil {
		return model.InvoiceResponse{}, err
	}
	if err := iu.ir.UpdateInvoice(context.Background(), &invoice, invoiceId); err != nil {
		fmt.Println(err)
		return model.InvoiceResponse{}, err
	}

	resInvoice := model.InvoiceResponse{}
	resInvoice.ID = invoice.ID
	resInvoice.Amount = invoice.Amount
	resInvoice.Status = invoice.Status

	return resInvoice, nil
}

func (iu *invoiceUsecase) DeleteInvoice(invoiceId uuid.UUID) error {
	if err := iu.ir.DeleteInvoice(context.Background(), invoiceId); err != nil {
		return err
	}
	return nil
}
