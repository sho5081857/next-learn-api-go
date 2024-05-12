package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Invoice struct {
	bun.BaseModel `bun:"invoices,alias:i"`

	ID         uuid.UUID `json:"id" bun:"type:char(36),default:uuid(),pk"`
	Amount     int       `json:"amount" bun:",notnull"`
	Status     string    `json:"status" bun:",notnull"`
	Date       time.Time `json:"date" bun:",nullzero,notnull"`
	Customer   Customer  `json:"customer" bun:"rel:belongs-to,join:customer_id=id"`
	CustomerId uuid.UUID `json:"customer_id" bun:"type:char(36),default:uuid()"`
}

type GetLatestInvoicesResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	ImageUrl string    `json:"image_url"`
	Email    string    `json:"email"`
	Amount   int       `json:"amount"`
}

type GetFilteredInvoicesResponse struct {
	ID         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	ImageUrl   string    `json:"image_url"`
	Amount     int       `json:"amount"`
	Date       time.Time `json:"date"`
	Status     string    `json:"status"`
}

type GetInvoiceByIdResponse struct {
	ID         uuid.UUID `json:"id"`
	CustomerId uuid.UUID `json:"customer_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
}

type InvoiceResponse struct {
	ID       uuid.UUID `json:"id"`
	Amount   int       `json:"amount"`
	Date     time.Time `json:"date"`
	Status   string    `json:"status"`
	Customer struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		ImageUrl string `json:"image_url"`
	} `json:"customer"`
}
