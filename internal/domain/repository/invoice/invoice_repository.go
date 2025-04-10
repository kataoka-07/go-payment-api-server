package invoice

import (
	"context"
	"go-payment-api-server/internal/domain/model"
)

type InvoiceRepository interface {
	Create(ctx context.Context, invoice *model.Invoice) error
}
