package invoice

import (
	"go-payment-api-server/internal/domain/errors"
	"go-payment-api-server/internal/domain/model"
	"go-payment-api-server/pkg/enum"
	"go-payment-api-server/pkg/util"

	"github.com/samber/lo"
)

type CreateInvoiceRequest struct {
	PartnerID     uint   `json:"partner_id"`
	IssueDate     string `json:"issue_date"` // yyyy-mm-dd
	DueDate       string `json:"due_date"`   // yyyy-mm-dd
	PaymentAmount int64  `json:"payment_amount"`
}

func (r *CreateInvoiceRequest) ToInvoice() (*model.Invoice, error) {
	return &model.Invoice{
		PartnerID:     int64(r.PartnerID),
		IssueDate:     lo.Must(util.ParseYMD(r.IssueDate)),
		DueDate:       lo.Must(util.ParseYMD(r.DueDate)),
		PaymentAmount: r.PaymentAmount,
		Status:        lo.ToPtr(enum.InvoiceStatusPending.String()),
	}, nil
}

func (r *CreateInvoiceRequest) Validate() error {
	if r.PartnerID == 0 || r.PaymentAmount <= 0 {
		return errors.ErrInvalidInvoice
	}
	return nil
}
