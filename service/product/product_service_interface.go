package product

import (
	"context"
	"mime/multipart"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type ProductServiceInterface interface {
	CreateProduct(ctx context.Context, product model.ProductModel, photos []*multipart.FileHeader, userID int) error
	UpdateProductByID(ctx context.Context, product model.ProductModel, photos []*multipart.FileHeader, id, userID int) error
}
