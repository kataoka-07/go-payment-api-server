package di

import (
	"go-payment-api-server/internal/infrastructure/mysql"
	invoicerepo "go-payment-api-server/internal/infrastructure/repository/invoice"
	userrepo "go-payment-api-server/internal/infrastructure/repository/user"
	invoicehdl "go-payment-api-server/internal/interface/handler/invoice"
	"go-payment-api-server/internal/interface/middleware"
	invoiceuc "go-payment-api-server/internal/usecase/invoice"
	"net/http"

	"github.com/go-chi/chi"
)

func InitializeRouter() http.Handler {
	cfg := mysql.NewDBConfigFromEnv()
	db := mysql.NewDB(cfg)

	r := chi.NewRouter()

	userRepo := userrepo.NewUserRepository(db)
	authMw := middleware.NewAuthMiddleware(userRepo)

	r.Use(authMw.Middleware)

	invoiceRepo := invoicerepo.NewInvoiceRepository(db)
	createInvoiceUC := invoiceuc.NewCreateInvoiceUseCase(invoiceRepo)
	invoicehdl.NewInvoiceHandler(r, createInvoiceUC)

	return r
}
