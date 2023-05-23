package detail_trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewDetailTrxRepository(DB *gorm.DB) DetailTrxRepositoryInterface {
	return &detailTrxRepository{
		DB: DB,
	}
}

type detailTrxRepository struct {
	*gorm.DB
}

func (detailTrxRepository *detailTrxRepository) BulkInsert(ctx context.Context, tx *gorm.DB, detailTrx []entity.DetailTrx) error {
	err := tx.Create(&detailTrx).Error
	if err != nil {
		return err
	}

	return nil
}
