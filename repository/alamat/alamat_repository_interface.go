package alamat

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type AlamatRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, alamat entity.Alamat) error
}
