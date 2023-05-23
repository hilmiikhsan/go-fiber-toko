package log_product

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewLogProductRepositoryInterface(DB *gorm.DB) LogProductRepositoryInterface {
	return &logProductRepository{
		DB: DB,
	}
}

type logProductRepository struct {
	*gorm.DB
}

func (logProductRepository *logProductRepository) Insert(ctx context.Context, tx *gorm.DB, logProduct entity.LogProduk) (entity.LogProduk, error) {
	err := tx.Create(&logProduct).Error
	if err != nil {
		return entity.LogProduk{}, err
	}

	return logProduct, nil
}
