package category

import (
	"context"
	"errors"

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

func (categoryRepository *categoryRepository) Update(ctx context.Context, tx *gorm.DB, category entity.Category, id int) error {
	err := tx.WithContext(ctx).Where("category.id = ?", id).Updates(&category).Error
	if err != nil {
		return err
	}

	return nil
}

func (categoryRepository *categoryRepository) FindByID(ctx context.Context, id int) (entity.Category, error) {
	category := entity.Category{}
	result := categoryRepository.DB.WithContext(ctx).Where("category.id = ?", id).Find(&category)
	if result.RowsAffected == 0 {
		return entity.Category{}, errors.New("Category not found")
	}

	return category, nil
}

func (categoryRepository *categoryRepository) Delete(ctx context.Context, tx *gorm.DB, category entity.Category, id int) error {
	err := tx.WithContext(ctx).Where("category.id = ?", id).Delete(&category).Error
	if err != nil {
		return err
	}

	return nil
}
