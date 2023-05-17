package province

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type ProvinceServiceInterface interface {
	GetProvinceCity(ctx context.Context) ([]model.Provinsi, error)
	GetProvinceDetail(ctx context.Context, provID string) (model.Provinsi, error)
}
