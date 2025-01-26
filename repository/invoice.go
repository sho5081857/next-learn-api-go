package repository

import (
	"context"
	"fmt"
	"next-learn-go/entity"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type InvoiceRepository interface {
	GetLatestInvoices(ctx context.Context, invoices *[]entity.Invoice, offset, limit int) error
	GetFilteredInvoices(ctx context.Context, invoices *[]entity.Invoice, query string, offset, limit int) error
	GetInvoiceCount(ctx context.Context) (int, error)
	GetInvoiceStatusCount(ctx context.Context) (int, int, error)
	GetInvoicesPages(ctx context.Context, query string, offset, limit int) (int, error)
	GetInvoiceById(ctx context.Context, invoice *entity.Invoice, invoiceId uuid.UUID) error
	CreateInvoice(ctx context.Context, invoice *entity.Invoice) error
	UpdateInvoice(ctx context.Context, invoice *entity.Invoice, invoiceId uuid.UUID) error
	DeleteInvoice(ctx context.Context, invoiceId uuid.UUID) error
}

type invoiceRepository struct {
	db *bun.DB
}

func NewInvoiceRepository(db *bun.DB) InvoiceRepository {
	return &invoiceRepository{db}
}

func (ir *invoiceRepository) GetLatestInvoices(ctx context.Context, invoices *[]entity.Invoice, offset, limit int) error {
	if err := ir.db.NewSelect().
		Model(invoices).
		Relation("Customer").
		Offset(offset).
		Limit(limit).
		OrderExpr("date").
		Scan(ctx); err != nil {
		return err
	}
	return nil
}

func (ir *invoiceRepository) GetInvoiceCount(ctx context.Context) (int, error) {
	count, err := ir.db.NewSelect().Model((*entity.Invoice)(nil)).Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ir *invoiceRepository) GetInvoiceStatusCount(ctx context.Context) (int, int, error) {
	pending, err := ir.db.NewSelect().Model((*entity.Invoice)(nil)).Where("status=?", "pending").Count(ctx)
	if err != nil {
		return 0, 0, err
	}
	paid, err := ir.db.NewSelect().Model((*entity.Invoice)(nil)).Where("status=?", "paid").Count(ctx)
	if err != nil {
		return 0, 0, err
	}
	return pending, paid, nil
}

func (ir *invoiceRepository) GetInvoicesPages(ctx context.Context, query string, offset, limit int) (int, error) {
	query = "%" + query + "%"
	count, err := ir.db.NewSelect().
		Model((*entity.Invoice)(nil)).
		Relation("Customer").
		WhereGroup("AND", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.WhereOr("Customer.name ILIKE ?", query).
				WhereOr("Customer.email ILIKE ?", query).
				WhereOr("amount::text ILIKE ?", query).
				WhereOr("date::text ILIKE ?", query).
				WhereOr("status ILIKE ?", query)
		}).
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (ir *invoiceRepository) GetFilteredInvoices(ctx context.Context, invoices *[]entity.Invoice, query string, offset, limit int) error {
	query = "%" + query + "%"
	if err := ir.db.NewSelect().
		Model(invoices).
		Relation("Customer").
		WhereGroup("AND", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.WhereOr("Customer.name ILIKE ?", query).
				WhereOr("Customer.email ILIKE ?", query).
				WhereOr("amount::text ILIKE ?", query).
				WhereOr("date::text ILIKE ?", query).
				WhereOr("status ILIKE ?", query)
		}).
		OrderExpr("date DESC").
		Limit(limit).
		Offset(offset).
		Scan(ctx); err != nil {
		return err
	}
	return nil
}

func (ir *invoiceRepository) GetInvoiceById(ctx context.Context, invoice *entity.Invoice, invoiceId uuid.UUID) error {
	if err := ir.db.NewSelect().
		Model(invoice).
		Relation("Customer").
		Where("i.id=?", invoiceId).
		Scan(ctx); err != nil {
		return err
	}
	return nil
}

func (ir *invoiceRepository) CreateInvoice(ctx context.Context, invoice *entity.Invoice) error {
	if _, err := ir.db.NewInsert().Model(invoice).Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (ir *invoiceRepository) UpdateInvoice(ctx context.Context, invoice *entity.Invoice, invoiceId uuid.UUID) error {
	result, err := ir.db.NewUpdate().
		Model(invoice).
		Column("customer_id", "amount", "status").
		Where("id=?", invoiceId).
		Exec(ctx)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}

func (ir *invoiceRepository) DeleteInvoice(ctx context.Context, invoiceId uuid.UUID) error {
	result, err := ir.db.NewDelete().
		Model(&entity.Invoice{}).
		Where("id=?", invoiceId).
		Exec(ctx)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}
	return nil
}
