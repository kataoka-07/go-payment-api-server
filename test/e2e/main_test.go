package e2e

import (
	"go-payment-api-server/pkg/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if logger.Log == nil {
		logger.Init()
	}

	code := m.Run()
	os.Exit(code)
}
