package trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewProductRepositoryInterface(DB *gorm.DB) TrxRepositoryInterface {
	return &trxRepository{
		DB: DB,
	}
}

type trxRepository struct {
	*gorm.DB
}

func (trxRepository *trxRepository) Insert(ctx context.Context, tx *gorm.DB, trx entity.Trx) (int, error) {
	err := tx.WithContext(ctx).Create(&trx).Error
	if err != nil {
		return 0, err
	}

	return trx.ID, nil
}
