package invoice

import (
	"go-payment-api-server/internal/domain/model"
	"go-payment-api-server/pkg/util"

	"github.com/samber/lo"
)

type InvoiceResponse struct {
	ID          int64  `json:"id"`
	CompanyID   int64  `json:"company_id"`
	PartnerID   int64  `json:"partner_id"`
	IssueDate   string `json:"issue_date"`
	DueDate     string `json:"due_date"`
	Amount      int64  `json:"amount"`
	Fee         int64  `json:"fee"`
	Tax         int64  `json:"tax"`
	TotalAmount int64  `json:"total_amount"`
	Status      string `json:"status"`
}

func FromModel(inv *model.Invoice) *InvoiceResponse {
	return &InvoiceResponse{
		ID:          inv.ID,
		CompanyID:   inv.CompanyID,
		PartnerID:   inv.PartnerID,
		IssueDate:   util.FormatYMD(inv.IssueDate),
		DueDate:     util.FormatYMD(inv.DueDate),
		Amount:      inv.PaymentAmount,
		Fee:         inv.Fee,
		Tax:         inv.Tax,
		TotalAmount: inv.TotalAmount,
		Status:      *inv.Status,
	}
}

func FromModels(invs []*model.Invoice) []*InvoiceResponse {
	return lo.Map(invs, func(inv *model.Invoice, _ int) *InvoiceResponse {
		return FromModel(inv)
	})
}
