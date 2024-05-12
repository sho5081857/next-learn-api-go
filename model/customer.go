package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Customer struct {
	bun.BaseModel `bun:"customers,alias:c"`

	ID            uuid.UUID `json:"id" bun:"type:char(36),default:uuid(),pk"`
	Name          string    `json:"name" bun:",notnull,type:varchar(45)"`
	Email         string    `json:"email" bun:",notnull,type:varchar(255)"`
	ImageUrl      string    `json:"image_url" bun:"type:varchar(255)"`
	Invoices      []Invoice `bun:"rel:has-many,join:id=customer_id"`
	TotalInvoices uint      `json:"total_invoices" bun:",scanonly"`
	TotalPending  uint      `json:"total_pending" bun:",scanonly"`
	TotalPaid     uint      `json:"total_paid" bun:",scanonly"`
}

type GetAllCustomerResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type GetFilteredCustomerResponse struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	ImageUrl      string    `json:"image_url"`
	TotalInvoices uint      `json:"total_invoices"`
	TotalPending  uint      `json:"total_pending"`
	TotalPaid     uint      `json:"total_paid"`
}
