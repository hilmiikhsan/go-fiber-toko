package alamat

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/alamat"
	"gorm.io/gorm"
)

func NewAlamatServiceInterface(alamatRepository *alamat.AlamatRepositoryInterface, db *gorm.DB) AlamatServiceInterface {
	return &alamatService{
		AlamatRepositoryInterface: *alamatRepository,
		DB:                        db,
	}
}

type alamatService struct {
	alamat.AlamatRepositoryInterface
	*gorm.DB
}

func (alamatService *alamatService) CreateAlamat(ctx context.Context, alamat model.AlamatModel, userID int) error {
	tx := alamatService.DB.Begin()

	alamatModel := entity.Alamat{
		IdUser:       userID,
		JudulAlamat:  alamat.JudulAlamat,
		NamaPenerima: alamat.NamaPenerima,
		NoTelp:       alamat.NoTelp,
		DetailAlamat: alamat.DetailAlamat,
	}

	err := alamatService.AlamatRepositoryInterface.Insert(ctx, tx, alamatModel)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (alamatService *alamatService) GetAllAlamat(ctx context.Context, params *struct{ model.ParamsModel }, userID int) ([]model.GetAlamatModel, error) {
	tmpAlamatData := []model.GetAlamatModel{}

	data, err := alamatService.AlamatRepositoryInterface.FindAll(ctx, params, userID)
	if err != nil {
		return tmpAlamatData, err
	}

	for _, x := range data {
		tmpAlamatData = append(tmpAlamatData, model.GetAlamatModel{
			ID:           x.ID,
			JudulAlamat:  x.JudulAlamat,
			NamaPenerima: x.NamaPenerima,
			NoTelp:       x.NoTelp,
			DetailAlamat: x.DetailAlamat,
		})
	}

	return tmpAlamatData, nil
}

func (alamatService *alamatService) GetAlamatByID(ctx context.Context, id, userID int) (model.GetAlamatModel, error) {
	tmpAlamatData := model.GetAlamatModel{}

	data, err := alamatService.AlamatRepositoryInterface.FindByID(ctx, id, userID)
	if err != nil {
		return tmpAlamatData, err
	}

	tmpAlamatData = model.GetAlamatModel{
		ID:           data.ID,
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		NoTelp:       data.NoTelp,
		DetailAlamat: data.DetailAlamat,
	}

	return tmpAlamatData, nil
}
