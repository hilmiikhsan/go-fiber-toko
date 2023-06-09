package trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type TrxServiceInterface interface {
	CreateTrx(ctx context.Context, trx model.TrxModel, userID int) error
	GetAllTrx(ctx context.Context, params *struct{ model.ParamsTrxModel }, userID int) ([]model.GetTrxModel, error)
	GetTrxByID(ctx context.Context, id, userID int) (model.GetTrxModel, error)
}
