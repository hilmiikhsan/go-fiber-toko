package alamat

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"gorm.io/gorm"
)

type AlamatRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, alamat entity.Alamat) error
	FindAll(ctx context.Context, params *struct{ model.ParamsModel }, userID int) ([]entity.Alamat, error)
	FindByID(ctx context.Context, id, userID int) (entity.Alamat, error)
	Update(ctx context.Context, tx *gorm.DB, alamat entity.Alamat, id, userID int) error
	Delete(ctx context.Context, tx *gorm.DB, alamat entity.Alamat, id, userID int) error
}
