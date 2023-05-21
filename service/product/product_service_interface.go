package product

import (
	"context"
	"mime/multipart"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type ProductServiceInterface interface {
	CreateProduct(ctx context.Context, product model.CreateProductModel, photos []*multipart.FileHeader, userID int) error
}
