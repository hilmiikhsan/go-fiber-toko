package toko

import (
	"context"
	"errors"

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

func (tokoRepository *tokoRepository) FindByID(ctx context.Context, id int) (entity.Toko, error) {
	toko := entity.Toko{}
	result := tokoRepository.DB.WithContext(ctx).Where("toko.id_user = ?", id).Find(&toko)
	if result.RowsAffected == 0 {
		return entity.Toko{}, errors.New("record not found")
	}

	return toko, nil
}
