package invoice

import (
	"context"
	"go-payment-api-server/internal/domain/model"
	invoicerepo "go-payment-api-server/internal/domain/repository/invoice"
	"go-payment-api-server/internal/infrastructure/query"
	"time"
)

type invoiceRepository struct {
	q *query.Query
}

func NewInvoiceRepository(q *query.Query) invoicerepo.InvoiceRepository {
	return &invoiceRepository{q: q}
}

func (r *invoiceRepository) Create(ctx context.Context, invoice *model.Invoice) error {
	return r.q.Invoice.WithContext(ctx).Create(invoice)
}

func (r *invoiceRepository) FindByDueDateRange(ctx context.Context, from, to time.Time, limit, offset int) ([]*model.Invoice, error) {
	return r.q.Invoice.WithContext(ctx).
		Where(r.q.Invoice.DueDate.Gte(from), r.q.Invoice.DueDate.Lte(to)).
		Limit(limit).
		Offset(offset).
		Find()
}
