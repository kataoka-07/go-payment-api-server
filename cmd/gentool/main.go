package main

import (
	"go-payment-api-server/internal/infrastructure/generator"
	"go-payment-api-server/pkg/logger"

	"github.com/joho/godotenv"
)

func main() {
	logger.Init()

	if err := godotenv.Load(); err != nil {
		logger.Log.Error("Error loading .env", "gen-err", err)
	}

	cfg := generator.NewGenConfigFromEnv()
	generator.GenerateModel(cfg)

	logger.Log.Info("Model generation complete")
}
