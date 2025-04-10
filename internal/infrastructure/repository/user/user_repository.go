package user

import (
	"context"
	"go-payment-api-server/internal/domain/model"
	"go-payment-api-server/internal/domain/repository/user"
	"go-payment-api-server/internal/infrastructure/query"
)

type userRepository struct {
	q *query.Query
}

func NewUserRepository(q *query.Query) user.UserRepository {
	return &userRepository{q: q}
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	return r.q.User.WithContext(ctx).Where(r.q.User.ID.Eq(id)).First()
}
