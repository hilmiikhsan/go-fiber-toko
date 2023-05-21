package foto_produk

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type FotoProdukRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, fotoProduk entity.FotoProduk) error
}
