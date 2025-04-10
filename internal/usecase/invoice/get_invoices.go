package invoice

import (
	"context"
	"go-payment-api-server/internal/domain/model"
	invoicerepo "go-payment-api-server/internal/domain/repository/invoice"
	"time"
)

type GetInvoicesUseCase interface {
	Execute(ctx context.Context, from, to time.Time, limit, offset int) ([]*model.Invoice, error)
}

type getInvoicesUseCase struct {
	invoiceRepo invoicerepo.InvoiceRepository
}

func NewGetInvoicesUseCase(invoiceRepo invoicerepo.InvoiceRepository) GetInvoicesUseCase {
	return &getInvoicesUseCase{invoiceRepo: invoiceRepo}
}

func (u *getInvoicesUseCase) Execute(ctx context.Context, from, to time.Time, limit, offset int) ([]*model.Invoice, error) {
	return u.invoiceRepo.FindByDueDateRange(ctx, from, to, limit, offset)
}
