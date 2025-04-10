package main

import (
	"go-payment-api-server/internal/di"
	"go-payment-api-server/pkg/logger"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/samber/lo"
)

func main() {
	logger.Init()
	_ = godotenv.Load(".env")
	router := di.InitializeRouter()

	envPort := os.Getenv("APP_PORT")
	port := lo.Ternary(envPort == "", "3001", envPort)

	logger.Log.Info("server running")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Log.Error("failed to start server", "server-err", err)
	}
}
