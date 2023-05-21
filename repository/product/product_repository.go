package product

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewProductRepositoryInterface(DB *gorm.DB) ProductRepositoryInterface {
	return &productRepository{
		DB: DB,
	}
}

type productRepository struct {
	*gorm.DB
}

func (productRepository *productRepository) Insert(ctx context.Context, tx *gorm.DB, product entity.Produk) (int, error) {
	err := tx.WithContext(ctx).Create(&product).Error
	if err != nil {
		return 0, err
	}

	return product.ID, nil
}

func (productRepository *productRepository) Update(ctx context.Context, tx *gorm.DB, product entity.Produk, id, idToko int) error {
	err := tx.WithContext(ctx).Where("produk.id = ?", id).Where("produk.id_toko = ?", idToko).Updates(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (productRepository *productRepository) FindByID(ctx context.Context, id int) (entity.Produk, error) {
	product := entity.Produk{}
	result := productRepository.DB.WithContext(ctx).Where("produk.id = ?", id).Find(&product)
	if result.RowsAffected == 0 {
		return entity.Produk{}, errors.New("Produk tidak ditemukan")
	}

	return product, nil
}
