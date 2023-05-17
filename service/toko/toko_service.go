package toko

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/hilmiikhsan/go_rest_api/entity"
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

	data, err := tokoService.TokoRepositoryInterface.FindByUserID(ctx, userID)
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

func (tokoService *tokoService) UpdateToko(ctx context.Context, namaToko string, photo *multipart.FileHeader, id, userID int) error {
	var photoData string

	if namaToko == "" {
		return errors.New("Nama toko is required")
	}

	data, err := tokoService.TokoRepositoryInterface.FindByIdAndUserID(ctx, id, userID)
	if err != nil {
		return err
	}

	if len(photo.Filename) > 0 {
		photoData = fmt.Sprintf("%d-%d%s", userID, time.Now().UnixNano(), filepath.Ext(photo.Filename))
		err := SaveFile(photo, userID)
		if err != nil {
			return err
		}
	}

	tx := tokoService.DB.Begin()

	tokoModel := entity.Toko{
		NamaToko: namaToko,
		UrlFoto:  photoData,
	}

	err = tokoService.TokoRepositoryInterface.Update(ctx, tx, tokoModel, data.ID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (tokoService *tokoService) GetAllToko(ctx context.Context, params *struct{ model.ParamsTokoModel }) ([]model.GetTokoModel, error) {
	response := []model.GetTokoModel{}

	if params.Page < 1 {
		params.Page = 1
	}

	if params.Limit < 1 {
		params.Limit = 10
	}

	data, err := tokoService.FindAll(ctx, params)
	if err != nil {
		return response, err
	}

	for _, x := range data {
		response = append(response, model.GetTokoModel{
			ID:       x.ID,
			NamaToko: x.NamaToko,
			UrlFoto:  x.UrlFoto,
		})
	}

	return response, nil
}

func SaveFile(file *multipart.FileHeader, userID int) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	photoData := fmt.Sprintf("%d-%d%s", userID, time.Now().UnixNano(), filepath.Ext(file.Filename))

	dst, err := os.Create("temp/" + photoData)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = src.Seek(0, 0); err != nil {
		return err
	}

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}

func (tokoService *tokoService) GeTokoByID(ctx context.Context, id int) (model.GetTokoModel, error) {
	response := model.GetTokoModel{}

	data, err := tokoService.TokoRepositoryInterface.FindByID(ctx, id)
	if err != nil {
		return response, err
	}

	response = model.GetTokoModel{
		ID:       data.ID,
		NamaToko: data.NamaToko,
		UrlFoto:  data.UrlFoto,
	}

	return response, nil
}
