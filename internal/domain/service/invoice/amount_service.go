package invoice

import "go-payment-api-server/internal/domain/model"

func CalculateAndFillAmounts(inv *model.Invoice) {
	const (
		feeRate = 0.04 // 手数料
		taxRate = 0.10 // 消費税
	)

	inv.FeeRate = feeRate
	inv.TaxRate = taxRate

	inv.Fee = int64(float64(inv.PaymentAmount) * feeRate)
	inv.Tax = int64(float64(inv.Fee) * taxRate)
	inv.TotalAmount = inv.PaymentAmount + inv.Fee + inv.Tax
}
