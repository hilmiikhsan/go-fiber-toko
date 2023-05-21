package foto_produk

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewFotoProdukRepositoryInterface(DB *gorm.DB) FotoProdukRepositoryInterface {
	return &fotoProdukRepository{
		DB: DB,
	}
}

type fotoProdukRepository struct {
	*gorm.DB
}

func (fotoProdukRepository *fotoProdukRepository) Insert(ctx context.Context, tx *gorm.DB, fotoProduk entity.FotoProduk) error {
	err := tx.WithContext(ctx).Create(&fotoProduk).Error
	if err != nil {
		return err
	}

	return nil
}
