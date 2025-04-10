package enum

type InvoiceStatus int

const (
	InvoiceStatusPending InvoiceStatus = iota
	InvoiceStatusProcessing
	InvoiceStatusPaid
	InvoiceStatusError
)

func (s InvoiceStatus) String() string {
	switch s {
	case InvoiceStatusPending:
		return "pending"
	case InvoiceStatusProcessing:
		return "processing"
	case InvoiceStatusPaid:
		return "paid"
	case InvoiceStatusError:
		return "error"
	default:
		return "pending"
	}
}
