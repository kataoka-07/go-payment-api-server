package di

import (
	"go-payment-api-server/internal/infrastructure/mysql"
	"go-payment-api-server/pkg/logger"
	"net/http"

	"github.com/go-chi/chi"
)

func InitializeRouter() http.Handler {
	cfg := mysql.NewDBConfigFromEnv()
	db := mysql.NewDB(cfg)

	// TODO: usecase/handler 初期化
	r := chi.NewRouter()

	// TODO: 後にhandler作成
	logger.Log.Info(db.Name())

	return r
}
