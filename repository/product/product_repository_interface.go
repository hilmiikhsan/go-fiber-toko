package product

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, product entity.Produk) (int, error)
}
