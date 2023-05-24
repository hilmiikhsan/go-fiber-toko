package log_product

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type LogProductRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, logProduct entity.LogProduk) (entity.LogProduk, error)
	FindByIdLogProduct(ctx context.Context, idLogProduct int) (entity.LogProduk, error)
	FindByIDsLogProduct(ctx context.Context, idsLogProduct []int) ([]entity.LogProduk, error)
}
