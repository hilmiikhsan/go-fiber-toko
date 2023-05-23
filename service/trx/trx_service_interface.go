package trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type TrxServiceInterface interface {
	CreateTrx(ctx context.Context, trx model.TrxModel, userID int) error
}
