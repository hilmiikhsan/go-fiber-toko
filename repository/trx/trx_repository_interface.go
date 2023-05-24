package trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"gorm.io/gorm"
)

type TrxRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, trx entity.Trx) (int, error)
	FindAll(ctx context.Context, params *struct{ model.ParamsTrxModel }, userID int) ([]entity.Trx, error)
	FindByID(ctx context.Context, id, userID int) (entity.Trx, error)
}
