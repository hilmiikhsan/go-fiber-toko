package category

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, tx *gorm.DB, category entity.Category) error
	Update(ctx context.Context, tx *gorm.DB, category entity.Category, id int) error
	FindByID(ctx context.Context, id int) (entity.Category, error)
	Delete(ctx context.Context, tx *gorm.DB, category entity.Category, id int) error
}
