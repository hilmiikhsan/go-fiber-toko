package product

import (
	"context"

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
