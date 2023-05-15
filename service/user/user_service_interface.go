package user

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type UserServiceInterface interface {
	GetProfile(ctx context.Context, email string) (model.User, error)
}
