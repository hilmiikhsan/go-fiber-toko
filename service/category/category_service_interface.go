package category

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type CategoryServiceInterface interface {
	CreateCategory(ctx context.Context, category model.CategoryModel, userID int) error
}
