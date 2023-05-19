package category

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, category entity.Category) error
}
