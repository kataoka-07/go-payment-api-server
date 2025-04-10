package invoice

import (
	"context"
	"go-payment-api-server/internal/domain/model"
	"time"
)

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *model.Invoice) error
	FindByDueDateRange(ctx context.Context, from, to time.Time, limit, offset int) ([]*model.Invoice, error)
}
