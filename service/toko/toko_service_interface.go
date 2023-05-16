package toko

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type TokoServiceInterface interface {
	GetMyToko(ctx context.Context, userID int) (model.TokoModel, error)
}
