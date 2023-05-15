package auth

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, user model.AuthRegisterModel) (string, error)
	Login(ctx context.Context, user model.AuthLoginModel) (model.AuthResponseModel, error)
}
