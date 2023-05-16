package toko

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewTokoRepositoryInterface(DB *gorm.DB) TokoRepositoryInterface {
	return &tokoRepository{
		DB: DB,
	}
}

type tokoRepository struct {
	*gorm.DB
}

func (tokoRepository *tokoRepository) Insert(ctx context.Context, tx *gorm.DB, toko entity.Toko) error {
	err := tx.WithContext(ctx).Create(&toko).Error
	if err != nil {
		return err
	}

	return nil
}
