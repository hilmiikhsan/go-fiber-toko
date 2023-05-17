package toko

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"gorm.io/gorm"
)

type TokoRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, toko entity.Toko) error
	FindByID(ctx context.Context, id int) (entity.Toko, error)
	Update(ctx context.Context, tx *gorm.DB, toko entity.Toko, id, userID int) error
	FindByIdAndUserID(ctx context.Context, id, userID int) (entity.Toko, error)
	FindAll(ctx context.Context, params *struct{ model.ParamsTokoModel }) ([]entity.Toko, error)
}
