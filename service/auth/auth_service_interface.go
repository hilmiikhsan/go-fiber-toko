package auth

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type AuthServiceInterface interface {
	Register(ctx context.Context, model model.AuthModel) (string, error)
}
