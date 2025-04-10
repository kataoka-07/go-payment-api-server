package invoice

import (
	"encoding/json"
	ctmerrors "go-payment-api-server/internal/domain/errors"
	dto "go-payment-api-server/internal/interface/dto/invoice"
	errorhandler "go-payment-api-server/internal/interface/handler/errors"
	usecase "go-payment-api-server/internal/usecase/invoice"
	"go-payment-api-server/pkg/response"
	"net/http"

	"github.com/go-chi/chi"
)

type InvoiceHandler interface {
	CreateInvoice(http.ResponseWriter, *http.Request)
}

type invoiceHandler struct {
	usecase usecase.CreateInvoiceUseCase
}

func NewInvoiceHandler(r chi.Router, uc usecase.CreateInvoiceUseCase) InvoiceHandler {
	h := &invoiceHandler{usecase: uc}
	r.Post("/api/invoices", h.CreateInvoice)
	return h
}

func (h *invoiceHandler) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateInvoiceRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorhandler.HandleError(w, ctmerrors.ErrInvalidInvoice)
		return
	}

	if err := req.Validate(); err != nil {
		errorhandler.HandleError(w, ctmerrors.ErrInvalidInvoice)
		return
	}

	invoice, err := req.ToInvoice()
	if err != nil {
		errorhandler.HandleError(w, ctmerrors.ErrInvalidInvoice)
		return
	}

	created, err := h.usecase.Execute(r.Context(), invoice)
	if err != nil {
		errorhandler.HandleError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, dto.FromModel(created))
}
