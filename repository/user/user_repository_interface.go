package user

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
)

type UserRepositoryInterface interface {
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	RegisterUser(ctx context.Context, user entity.User) error
}
