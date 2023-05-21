package product

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, product entity.Produk) (int, error)
	Update(ctx context.Context, tx *gorm.DB, product entity.Produk, id, idToko int) error
	FindByID(ctx context.Context, id int) (entity.Produk, error)
	Delete(ctx context.Context, tx *gorm.DB, product entity.Produk, id, idToko int) error
}
