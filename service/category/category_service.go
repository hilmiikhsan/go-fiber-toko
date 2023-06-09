package category

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/category"
	"github.com/hilmiikhsan/go_rest_api/repository/user"
	"gorm.io/gorm"
)

func NewCategoryServiceInterface(categoryRepository *category.CategoryRepositoryInterface, db *gorm.DB, userRepository *user.UserRepositoryInterface) CategoryServiceInterface {
	return &categoryService{
		CategoryRepositoryInterface: *categoryRepository,
		DB:                          db,
		UserRepositoryInterface:     *userRepository,
	}
}

type categoryService struct {
	category.CategoryRepositoryInterface
	*gorm.DB
	user.UserRepositoryInterface
}

func (categoryService *categoryService) CreateCategory(ctx context.Context, category model.CategoryModel, userID int) error {
	userData, err := categoryService.UserRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if !userData.IsAdmin {
		return errors.New("Unauthorized")
	}

	tx := categoryService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	categoryModel := entity.Category{
		NamaCategory: category.NamaCategory,
	}

	err = categoryService.CategoryRepositoryInterface.Insert(ctx, tx, categoryModel)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (categoryService *categoryService) UpdateCategoryByID(ctx context.Context, id, userID int, category model.CategoryModel) error {
	userData, err := categoryService.UserRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if !userData.IsAdmin {
		return errors.New("Unauthorized")
	}

	categoryData, err := categoryService.CategoryRepositoryInterface.FindByID(ctx, id)
	if err != nil {
		return err
	}

	tx := categoryService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	categoryModel := entity.Category{
		NamaCategory: category.NamaCategory,
	}

	err = categoryService.CategoryRepositoryInterface.Update(ctx, tx, categoryModel, categoryData.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (categoryService *categoryService) DeleteCategoryByID(ctx context.Context, id, userID int) error {
	userData, err := categoryService.UserRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if !userData.IsAdmin {
		return errors.New("Unauthorized")
	}

	categoryData, err := categoryService.CategoryRepositoryInterface.FindByID(ctx, id)
	if err != nil {
		return err
	}

	tx := categoryService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = categoryService.CategoryRepositoryInterface.Delete(ctx, tx, categoryData, categoryData.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (categoryService *categoryService) GetAllCategory(ctx context.Context, userID int) ([]model.GetCategoryModel, error) {
	tmpCategoryData := []model.GetCategoryModel{}

	userData, err := categoryService.UserRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return tmpCategoryData, err
	}

	if !userData.IsAdmin {
		return tmpCategoryData, errors.New("Unauthorized")
	}

	data, err := categoryService.CategoryRepositoryInterface.FindAll(ctx)
	if err != nil {
		return tmpCategoryData, err
	}

	for _, x := range data {
		tmpCategoryData = append(tmpCategoryData, model.GetCategoryModel{
			ID:           x.ID,
			NamaCategory: x.NamaCategory,
		})
	}

	return tmpCategoryData, nil
}

func (categoryService *categoryService) GetCategoryByID(ctx context.Context, id, userID int) (model.GetCategoryModel, error) {
	tmpCategoryData := model.GetCategoryModel{}

	userData, err := categoryService.UserRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return tmpCategoryData, err
	}

	if !userData.IsAdmin {
		return tmpCategoryData, errors.New("Unauthorized")
	}

	data, err := categoryService.CategoryRepositoryInterface.FindByID(ctx, id)
	if err != nil {
		return tmpCategoryData, err
	}

	tmpCategoryData = model.GetCategoryModel{
		ID:           data.ID,
		NamaCategory: data.NamaCategory,
	}

	return tmpCategoryData, nil
}
