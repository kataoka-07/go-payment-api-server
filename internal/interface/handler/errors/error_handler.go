package errors

import (
	"errors"
	ctmerrors "go-payment-api-server/internal/domain/errors"
	"go-payment-api-server/pkg/logger"
	"go-payment-api-server/pkg/response"
	"net/http"
)

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func HandleError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError

	switch {
	case errors.Is(err, ctmerrors.ErrInvalidInvoice):
		status = http.StatusBadRequest
	case errors.Is(err, ctmerrors.ErrInvoiceCreation):
		status = http.StatusInternalServerError
	case errors.Is(err, ctmerrors.ErrUnauthorizedUser):
		status = http.StatusUnauthorized
	case errors.Is(err, ctmerrors.ErrUserNotFound):
		status = http.StatusNotFound
	case errors.Is(err, ctmerrors.ErrInvalidAccessToken):
		status = http.StatusUnauthorized
	}

	msg := ctmerrors.Message(err)

	logger.Log.Error("error occurred", "error", err)

	resp := APIError{Code: status, Message: msg}

	response.WriteJSON(w, status, resp)
}
