package alamat

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type AlamatServiceInterface interface {
	CreateAlamat(ctx context.Context, alamat model.AlamatModel, userID int) error
	GetAllAlamat(ctx context.Context, params *struct{ model.ParamsModel }, userID int) ([]model.GetAlamatModel, error)
	GetAlamatByID(ctx context.Context, id, userID int) (model.GetAlamatModel, error)
	UpdateAlamatByID(ctx context.Context, id, userID int, alamat model.UpdateAlamatModel) error
}
