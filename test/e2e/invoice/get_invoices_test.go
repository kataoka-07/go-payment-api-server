package invoice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"go-payment-api-server/internal/di"
	"go-payment-api-server/pkg/logger"
)

type getInvoicesTestCase struct {
	name         string
	queryParams  map[string]string
	headerUserID string
	wantStatus   int
	wantLength   int
	wantError    bool
}

func TestGetInvoices(t *testing.T) {
	logger.Init()

	tests := []getInvoicesTestCase{
		{
			name: "[正常系] 期間指定で請求書取得",
			queryParams: map[string]string{
				"from": "2025-04-28",
				"to":   "2025-04-30",
			},
			headerUserID: "1",
			wantStatus:   http.StatusOK,
			wantLength:   6,
			wantError:    false,
		},
		{
			name: "[正常系] limit指定で最大3件取得",
			queryParams: map[string]string{
				"from":  "2025-04-30",
				"to":    "2025-04-30",
				"limit": "3",
			},
			headerUserID: "1",
			wantStatus:   http.StatusOK,
			wantLength:   3,
		},
		{
			name: "[異常系] limitに負数",
			queryParams: map[string]string{
				"from":  "2025-04-10",
				"to":    "2025-04-30",
				"limit": "-1",
			},
			headerUserID: "1",
			wantStatus:   http.StatusOK,
			wantLength:   7,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(di.InitializeRouter(logger.Log))
			defer server.Close()

			baseURL := fmt.Sprintf("%s/api/invoices", server.URL)
			query := url.Values{}
			for k, v := range tt.queryParams {
				query.Set(k, v)
			}
			fullURL := fmt.Sprintf("%s?%s", baseURL, query.Encode())

			req, _ := http.NewRequest(http.MethodGet, fullURL, nil)
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("X-User-ID", tt.headerUserID)
			req.Header.Set("Authorization", "Bearer dummy")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("request failed: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.wantStatus {
				t.Errorf("expected status %d, got %d", tt.wantStatus, resp.StatusCode)
			}

			if !tt.wantError {
				var res []map[string]interface{}
				if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
					t.Fatalf("response parse error: %v", err)
				}

				if len(res) != tt.wantLength {
					t.Errorf("expected %d invoices, got %d", tt.wantLength, len(res))
				}
			}
		})
	}
}
