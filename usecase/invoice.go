package usecase

import (
	"context"
	"next-learn-go/entity"
	"next-learn-go/repository"
	"next-learn-go/validator"

	"github.com/google/uuid"
)

type InvoiceUseCase interface {
	GetLatestInvoices(offset, limit int) ([]entity.GetLatestInvoicesResponse, error)
	GetFilteredInvoices(query string, offset, limit int) ([]entity.GetFilteredInvoicesResponse, error)
	GetInvoiceCount() (int, error)
	GetInvoiceStatusCount() (int, int, error)
	GetInvoicesPages(query string, offset, limit int) (int, error)
	GetInvoiceById(invoiceId uuid.UUID) (entity.GetInvoiceByIdResponse, error)
	CreateInvoice(invoice entity.Invoice) (entity.InvoiceResponse, error)
	UpdateInvoice(invoice entity.Invoice, invoiceId uuid.UUID) (entity.InvoiceResponse, error)
	DeleteInvoice(invoiceId uuid.UUID) error
}

type invoiceUseCase struct {
	ir repository.InvoiceRepository
	iv validator.InvoiceValidator
}

func NewInvoiceUseCase(ir repository.InvoiceRepository, iv validator.InvoiceValidator) InvoiceUseCase {
	return &invoiceUseCase{ir, iv}
}

func (iu *invoiceUseCase) GetLatestInvoices(offset, limit int) ([]entity.GetLatestInvoicesResponse, error) {
	invoices := []entity.Invoice{}
	if err := iu.ir.GetLatestInvoices(context.Background(), &invoices, offset, limit); err != nil {
		return nil, err
	}
	resInvoices := []entity.GetLatestInvoicesResponse{}
	for _, v := range invoices {
		i := entity.GetLatestInvoicesResponse{}
		i.ID = v.ID
		i.Name = v.Customer.Name
		i.ImageUrl = v.Customer.ImageUrl
		i.Email = v.Customer.Email
		i.Amount = v.Amount
		resInvoices = append(resInvoices, i)
	}
	return resInvoices, nil
}

func (iu *invoiceUseCase) GetFilteredInvoices(query string, offset, limit int) ([]entity.GetFilteredInvoicesResponse, error) {
	invoices := []entity.Invoice{}
	if err := iu.ir.GetFilteredInvoices(context.Background(), &invoices, query, offset, limit); err != nil {
		return nil, err
	}
	resInvoices := []entity.GetFilteredInvoicesResponse{}
	for _, v := range invoices {
		i := entity.GetFilteredInvoicesResponse{}
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

func (iu *invoiceUseCase) GetInvoiceCount() (int, error) {
	count, err := iu.ir.GetInvoiceCount(context.Background())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iu *invoiceUseCase) GetInvoiceStatusCount() (int, int, error) {
	pending, paid, err := iu.ir.GetInvoiceStatusCount(context.Background())
	if err != nil {
		return 0, 0, err
	}
	return pending, paid, nil
}

func (iu *invoiceUseCase) GetInvoicesPages(query string, offset, limit int) (int, error) {
	count, err := iu.ir.GetInvoicesPages(context.Background(), query, offset, limit)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (iu *invoiceUseCase) GetInvoiceById(invoiceId uuid.UUID) (entity.GetInvoiceByIdResponse, error) {
	invoice := entity.Invoice{}
	if err := iu.ir.GetInvoiceById(context.Background(), &invoice, invoiceId); err != nil {
		return entity.GetInvoiceByIdResponse{}, err
	}

	resInvoice := entity.GetInvoiceByIdResponse{}
	resInvoice.ID = invoice.ID
	resInvoice.CustomerId = invoice.Customer.ID
	resInvoice.Amount = invoice.Amount
	resInvoice.Status = invoice.Status

	return resInvoice, nil
}

func (iu *invoiceUseCase) CreateInvoice(invoice entity.Invoice) (entity.InvoiceResponse, error) {
	if err := iu.iv.InvoiceValidate(invoice); err != nil {
		return entity.InvoiceResponse{}, err
	}
	if err := iu.ir.CreateInvoice(context.Background(), &invoice); err != nil {
		return entity.InvoiceResponse{}, err
	}

	resInvoice := entity.InvoiceResponse{}
	resInvoice.ID = invoice.ID
	resInvoice.Amount = invoice.Amount
	resInvoice.Date = invoice.Date
	resInvoice.Status = invoice.Status
	resInvoice.Customer.Name = invoice.Customer.Name
	resInvoice.Customer.Email = invoice.Customer.Email
	resInvoice.Customer.ImageUrl = invoice.Customer.ImageUrl

	return resInvoice, nil
}

func (iu *invoiceUseCase) UpdateInvoice(invoice entity.Invoice, invoiceId uuid.UUID) (entity.InvoiceResponse, error) {
	if err := iu.iv.InvoiceValidate(invoice); err != nil {
		return entity.InvoiceResponse{}, err
	}
	if err := iu.ir.UpdateInvoice(context.Background(), &invoice, invoiceId); err != nil {
		return entity.InvoiceResponse{}, err
	}

	resInvoice := entity.InvoiceResponse{}
	resInvoice.ID = invoice.ID
	resInvoice.Amount = invoice.Amount
	resInvoice.Status = invoice.Status

	return resInvoice, nil
}

func (iu *invoiceUseCase) DeleteInvoice(invoiceId uuid.UUID) error {
	if err := iu.ir.DeleteInvoice(context.Background(), invoiceId); err != nil {
		return err
	}
	return nil
}
