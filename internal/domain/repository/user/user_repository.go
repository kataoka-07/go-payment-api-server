package user

import (
	"context"
	"go-payment-api-server/internal/domain/model"
)

type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
}
