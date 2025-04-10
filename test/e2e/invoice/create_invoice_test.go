package invoice

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-payment-api-server/internal/di"
	"go-payment-api-server/pkg/logger"
)

type invoiceTestCase struct {
	name            string
	body            map[string]interface{}
	headerUserID    string
	wantStatus      int
	wantTotalAmount int64
	wantError       bool
}

func TestCreateInvoice(t *testing.T) {
	logger.Init()

	tests := []invoiceTestCase{
		{
			name: "[正常系] 正しい請求書データ",
			body: map[string]interface{}{
				"partner_id":     1,
				"issue_date":     "2025-04-10",
				"due_date":       "2025-05-10",
				"payment_amount": 10000,
			},
			headerUserID:    "1",
			wantStatus:      http.StatusOK,
			wantTotalAmount: 10440,
			wantError:       false,
		},
		{
			name: "[異常系] 不正な支払金額（0）",
			body: map[string]interface{}{
				"partner_id":     1,
				"issue_date":     "2025-04-10",
				"due_date":       "2025-05-10",
				"payment_amount": 0,
			},
			headerUserID: "1",
			wantStatus:   http.StatusBadRequest,
			wantError:    true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(di.InitializeRouter(logger.Log))
			defer server.Close()

			jsonBody, _ := json.Marshal(tt.body)
			req, _ := http.NewRequest(http.MethodPost, server.URL+"/api/invoices", bytes.NewBuffer(jsonBody))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-ID", tt.headerUserID)
			req.Header.Set("Authorization", "Bearer dummy")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("failed request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected %d, actually %d", tt.wantStatus, resp.StatusCode)
			}

			if !tt.wantError {
				var res struct {
					TotalAmount int64 `json:"total_amount"`
				}
				if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
					t.Fatalf("failed to parse response: %v", err)
				}

				if res.TotalAmount != tt.wantTotalAmount {
					t.Errorf("expected %d, actually %d", tt.wantTotalAmount, res.TotalAmount)
				}
			}
		})
	}
}
