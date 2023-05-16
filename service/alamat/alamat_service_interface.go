package alamat

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type AlamatServiceInterface interface {
	CreateAlamat(ctx context.Context, alamat model.AlamatModel, userID int) error
}
