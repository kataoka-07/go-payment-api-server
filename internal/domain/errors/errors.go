package errors

import "errors"

var (
	ErrInvalidInvoice     = errors.New("invalid invoice data")
	ErrInvoiceCreation    = errors.New("failed to create invoice")
	ErrUnauthorizedUser   = errors.New("unauthorized user")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrInvalidPeriod      = errors.New("invalid period")
)

func Message(err error) string {
	switch {
	case errors.Is(err, ErrInvalidInvoice):
		return "invalid invoice data"
	case errors.Is(err, ErrInvoiceCreation):
		return "failed to create invoice"
	case errors.Is(err, ErrUnauthorizedUser):
		return "unauthorized user"
	case errors.Is(err, ErrUserNotFound):
		return "user not found"
	case errors.Is(err, ErrInvalidAccessToken):
		return "invalid access token"
	case errors.Is(err, ErrInvalidPeriod):
		return "invalid period"
	default:
		return "internal server error"
	}
}
