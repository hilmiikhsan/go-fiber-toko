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

func (fotoProdukRepository *fotoProdukRepository) Update(ctx context.Context, tx *gorm.DB, fotoProduk entity.FotoProduk, idFoto, productID int) error {
	err := tx.WithContext(ctx).Where("foto_produk.id = ? AND foto_produk.id_produk = ?", idFoto, productID).Updates(&fotoProduk).Error
	if err != nil {
		return err
	}

	// err := tx.WithContext(ctx).Where("foto_produk.id_produk = ?", productID).Save(&fotoProduk).Error
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (fotoProdukRepository *fotoProdukRepository) FindByProductID(ctx context.Context, tx *gorm.DB, productID int) ([]entity.FotoProduk, error) {
	results := []entity.FotoProduk{}
	err := fotoProdukRepository.DB.WithContext(ctx).Where("foto_produk.id_produk = ?", productID).Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}
