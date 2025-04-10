package invoice

import (
	"context"
	"fmt"
	"go-payment-api-server/internal/domain/errors"
	ctmerrors "go-payment-api-server/internal/domain/errors"
	"go-payment-api-server/internal/domain/model"
	invoicerepo "go-payment-api-server/internal/domain/repository/invoice"
	domainservice "go-payment-api-server/internal/domain/service/invoice"
	"go-payment-api-server/pkg/contextkey"
)

type CreateInvoiceUseCase interface {
	Execute(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error)
}

type createInvoiceUseCase struct {
	invoiceRepo invoicerepo.InvoiceRepository
}

func NewCreateInvoiceUseCase(invoiceRepo invoicerepo.InvoiceRepository) CreateInvoiceUseCase {
	return &createInvoiceUseCase{invoiceRepo: invoiceRepo}
}

func (u *createInvoiceUseCase) Execute(ctx context.Context, invoice *model.Invoice) (*model.Invoice, error) {
	domainservice.CalculateAndFillAmounts(invoice)

	companyID, ok := ctx.Value(contextkey.ContextKeyCompanyID).(int64)
	if !ok {
		return nil, ctmerrors.ErrUnauthorizedUser
	}
	invoice.CompanyID = companyID

	if err := u.invoiceRepo.Create(ctx, invoice); err != nil {
		return nil, fmt.Errorf("%w: %v", errors.ErrInvoiceCreation, err)
	}
	return invoice, nil
}
