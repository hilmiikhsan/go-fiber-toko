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
}
