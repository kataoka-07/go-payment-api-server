package invoice

import (
	"context"
	"go-payment-api-server/internal/domain/model"
	invoicerepo "go-payment-api-server/internal/domain/repository/invoice"

	"gorm.io/gorm"
)

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) invoicerepo.InvoiceRepository {
	return &invoiceRepository{db: db}
}

func (r *invoiceRepository) Create(ctx context.Context, invoice *model.Invoice) error {
	return r.db.WithContext(ctx).Create(invoice).Error
}
