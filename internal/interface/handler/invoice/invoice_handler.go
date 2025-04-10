package invoice

import (
	"encoding/json"
	ctmerrors "go-payment-api-server/internal/domain/errors"
	dto "go-payment-api-server/internal/interface/dto/invoice"
	errorhandler "go-payment-api-server/internal/interface/handler/errors"
	usecase "go-payment-api-server/internal/usecase/invoice"
	"go-payment-api-server/pkg/response"
	"go-payment-api-server/pkg/util"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type InvoiceHandler interface {
	CreateInvoice(http.ResponseWriter, *http.Request)
	GetInvoices(http.ResponseWriter, *http.Request)
}

type invoiceHandler struct {
	createUsecase usecase.CreateInvoiceUseCase
	getUsecase    usecase.GetInvoicesUseCase
}

func NewInvoiceHandler(
	r chi.Router,
	createUC usecase.CreateInvoiceUseCase,
	getUC usecase.GetInvoicesUseCase,
) InvoiceHandler {
	h := &invoiceHandler{
		createUsecase: createUC,
		getUsecase:    getUC,
	}
	r.Post("/api/invoices", h.CreateInvoice)
	r.Get("/api/invoices", h.GetInvoices)
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

	created, err := h.createUsecase.Execute(r.Context(), invoice)
	if err != nil {
		errorhandler.HandleError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, dto.FromModel(created))
}

func (h *invoiceHandler) GetInvoices(w http.ResponseWriter, r *http.Request) {
	const defaultLimit = 10

	queryParams := r.URL.Query()
	fromStr := queryParams.Get("from")
	toStr := queryParams.Get("to")

	limit := defaultLimit
	if l := queryParams.Get("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	offset := 0
	if o := queryParams.Get("offset"); o != "" {
		if parsed, err := strconv.Atoi(o); err == nil && parsed >= 0 {
			offset = parsed
		}
	}

	from, err := util.ParseYMD(fromStr)
	if err != nil {
		errorhandler.HandleError(w, ctmerrors.ErrInvalidPeriod)
		return
	}

	to, err := util.ParseYMD(toStr)
	if err != nil {
		errorhandler.HandleError(w, ctmerrors.ErrInvalidPeriod)
		return
	}

	invoices, err := h.getUsecase.Execute(r.Context(), from, to, limit, offset)
	if err != nil {
		errorhandler.HandleError(w, err)
		return
	}

	response.WriteJSON(w, http.StatusOK, dto.FromModels(invoices))
}
