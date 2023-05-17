package toko

import (
	"context"
	"mime/multipart"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type TokoServiceInterface interface {
	GetMyToko(ctx context.Context, userID int) (model.TokoModel, error)
	// UpdateToko(ctx context.Context, toko model.UppdateTokoModel, id, userID int) error
	UpdateToko(ctx context.Context, namaToko string, photo *multipart.FileHeader, id, userID int) error
}
