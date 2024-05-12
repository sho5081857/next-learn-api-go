package validator

import (
	"next-learn-go/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IInvoiceValidator interface {
	InvoiceValidate(invoice model.Invoice) error
}

type invoiceValidator struct{}

func NewInvoiceValidator() IInvoiceValidator {
	return &invoiceValidator{}
}

func (tv *invoiceValidator) InvoiceValidate(invoice model.Invoice) error {
	return validation.ValidateStruct(&invoice,
		validation.Field(
			&invoice.CustomerId,
			validation.Required.Error("CustomerId is required"),
		),
		validation.Field(
			&invoice.Amount,
			validation.Required.Error("Amount is required"),
		),
		validation.Field(
			&invoice.Status,
			validation.Required.Error("Status is required"),
			validation.In(invoice.Status, "pending", "paid").Error("Status must be pending or paid"),
		),
	)
}
