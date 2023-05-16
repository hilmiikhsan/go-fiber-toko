package toko

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type TokoRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, toko entity.Toko) error
	FindByID(ctx context.Context, id int) (entity.Toko, error)
}
