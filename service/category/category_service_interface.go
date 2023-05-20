package category

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, category model.CategoryModel, userID int) error
	UpdateCategoryByID(ctx context.Context, id, userID int, category model.CategoryModel) error
	DeleteCategoryByID(ctx context.Context, id, userID int) error
	GetAllCategory(ctx context.Context, userID int) ([]model.GetCategoryModel, error)
	GetCategoryByID(ctx context.Context, id, userID int) (model.GetCategoryModel, error)
}
