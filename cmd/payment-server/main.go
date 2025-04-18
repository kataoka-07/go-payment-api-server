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
	router := di.InitializeRouter(logger.Log)

	envPort := os.Getenv("APP_PORT")
	port := lo.Ternary(envPort == "", "8000", envPort)

	logger.Log.Info("Server running")
	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Log.Error("Failed to start server", "server-err", err)
	}
}
