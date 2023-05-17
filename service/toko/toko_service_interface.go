package toko

import (
	"context"
	"mime/multipart"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type TokoServiceInterface interface {
	GetMyToko(ctx context.Context, userID int) (model.TokoModel, error)
	UpdateToko(ctx context.Context, namaToko string, photo *multipart.FileHeader, id, userID int) error
	GetAllToko(ctx context.Context, params *struct{ model.ParamsTokoModel }) ([]model.GetTokoModel, error)
	GeTokoByID(ctx context.Context, id int) (model.GetTokoModel, error)
}
