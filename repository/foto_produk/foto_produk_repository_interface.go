package foto_produk

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type FotoProdukRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, fotoProduk entity.FotoProduk) error
	Update(ctx context.Context, tx *gorm.DB, fotoProduk entity.FotoProduk, idFoto, productID int) error
	FindByProductID(ctx context.Context, tx *gorm.DB, productID int) ([]entity.FotoProduk, error)
	FindAll(ctx context.Context, idToko []int) ([]entity.FotoProduk, error)
}
