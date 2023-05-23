package trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type TrxRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, trx entity.Trx) (int, error)
}
