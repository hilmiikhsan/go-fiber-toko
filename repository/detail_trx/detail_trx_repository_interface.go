package detail_trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type DetailTrxRepositoryInterface interface {
	BulkInsert(ctx context.Context, tx *gorm.DB, detailTrx []entity.DetailTrx) error
	FindByIdTrx(ctx context.Context, idTrx int) ([]entity.DetailTrx, error)
	FindByIDsTrx(ctx context.Context, idsTrx []int) ([]entity.DetailTrx, error)
}
