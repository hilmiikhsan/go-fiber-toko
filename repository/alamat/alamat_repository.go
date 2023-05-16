package alamat

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewAlamatRepositoryInterface(DB *gorm.DB) AlamatRepositoryInterface {
	return &alamatRepository{
		DB: DB,
	}
}

type alamatRepository struct {
	*gorm.DB
}

func (alamatRepository *alamatRepository) Insert(ctx context.Context, tx *gorm.DB, alamat entity.Alamat) error {
	err := tx.WithContext(ctx).Create(&alamat).Error
	if err != nil {
		return err
	}

	return nil
}
