package repository

import (
	"context"
	"next-learn-go/model"

	"github.com/uptrace/bun"
)

type ICustomerRepository interface {
	GetAllCustomers(ctx context.Context, customers *[]model.Customer) error
	GetFilteredCustomers(ctx context.Context, customers *[]model.Customer, filter string) error
	GetCustomerCount(ctx context.Context) (int, error)
}

type customerRepository struct {
	db *bun.DB
}

func NewCustomerRepository(db *bun.DB) ICustomerRepository {
	return &customerRepository{db}
}

func (cr *customerRepository) GetAllCustomers(ctx context.Context, customers *[]model.Customer) error {
	if err := cr.db.NewSelect().
		Model(customers).
		Scan(ctx); err != nil {
		return err
	}
	return nil
}

func (cr *customerRepository) GetFilteredCustomers(ctx context.Context, customers *[]model.Customer, filter string) error {
	query := "%" + filter + "%"
	if err := cr.db.NewSelect().
		Model(customers).
		Column("id", "name", "email", "image_url").
		ColumnExpr("COUNT(invoices.id) AS total_invoices").
		ColumnExpr("SUM(CASE WHEN invoices.status = 'pending' THEN invoices.amount ELSE 0 END) AS total_pending").
		ColumnExpr("SUM(CASE WHEN invoices.status = 'paid' THEN invoices.amount ELSE 0 END) AS total_paid").
		Join("LEFT JOIN invoices ON c.id = invoices.customer_id").
		WhereGroup("AND", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.WhereOr("c.name ILIKE ?", query).
				WhereOr("c.email ILIKE ?", query)
		}).
		Group("c.id", "c.name", "c.email", "c.image_url").
		Order("c.name ASC").
		Scan(ctx); err != nil {
		return err
	}
	return nil
}
func (cr *customerRepository) GetCustomerCount(ctx context.Context) (int, error) {
	count, err := cr.db.NewSelect().Model((*model.Customer)(nil)).Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}
