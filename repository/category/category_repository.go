package category

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewCategoryRepositoryInterface(DB *gorm.DB) CategoryRepositoryInterface {
	return &categoryRepository{
		DB: DB,
	}
}

type categoryRepository struct {
	*gorm.DB
}

func (categoryRepository *categoryRepository) Insert(ctx context.Context, tx *gorm.DB, category entity.Category) error {
	err := tx.WithContext(ctx).Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}
