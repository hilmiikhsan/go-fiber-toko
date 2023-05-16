package toko

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/toko"
	"gorm.io/gorm"
)

func NewTokoServiceInterface(tokoRepository *toko.TokoRepositoryInterface, db *gorm.DB) TokoServiceInterface {
	return &tokoService{
		TokoRepositoryInterface: *tokoRepository,
		DB:                      db,
	}
}

type tokoService struct {
	toko.TokoRepositoryInterface
	*gorm.DB
}

func (tokoService *tokoService) GetMyToko(ctx context.Context, userID int) (model.TokoModel, error) {
	response := model.TokoModel{}

	data, err := tokoService.TokoRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return response, err
	}

	response = model.TokoModel{
		ID:       data.ID,
		NamaToko: data.NamaToko,
		UrlFoto:  data.UrlFoto,
		UserID:   data.IdUser,
	}

	return response, nil
}
