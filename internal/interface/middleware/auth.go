package middleware

import (
	"context"
	ctmerrors "go-payment-api-server/internal/domain/errors"
	"go-payment-api-server/internal/domain/repository/user"
	"go-payment-api-server/internal/interface/handler/errors"
	"go-payment-api-server/pkg/contextkey"
	"net/http"
	"strconv"
	"strings"
)

type AuthMiddleware struct {
	userRepo user.UserRepository
}

func NewAuthMiddleware(repo user.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{userRepo: repo}
}

func (m *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		token := extractBearerToken(r.Header.Get("Authorization"))
		if token == "" {
			errors.HandleError(w, ctmerrors.ErrInvalidAccessToken)
			return
		}

		// NOTE: トークンの有効期限確認（実際には JWT検証などに置き換え）
		if !isTokenValid(token) {
			errors.HandleError(w, ctmerrors.ErrInvalidAccessToken)
			return
		}

		userIDStr := r.Header.Get("X-User-ID")
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			errors.HandleError(w, ctmerrors.ErrUnauthorizedUser)
			return
		}

		user, err := m.userRepo.FindByID(ctx, userID)
		if err != nil {
			errors.HandleError(w, ctmerrors.ErrUserNotFound)
			return
		}

		ctx = context.WithValue(ctx, contextkey.ContextKeyCompanyID, user.CompanyID)
		ctx = context.WithValue(ctx, contextkey.ContextKeyUserID, user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractBearerToken(header string) string {
	if header == "" {
		return ""
	}
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	return parts[1]
}

func isTokenValid(_ string) bool {
	// NOTE: 有効期限は手動で判定を管理して簡易に処理
	expired := false
	return !expired
}
