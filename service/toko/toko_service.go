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

func (tokoService *tokoService) UpdateToko(ctx context.Context, namaToko string, photo *multipart.FileHeader, id, userID int) error {
	var photoData string

	if namaToko == "" {
		return errors.New("Nama toko is required")
	}

	if len(photo.Filename) > 0 {
		photoData = fmt.Sprintf("%d-%d%s", userID, time.Now().UnixNano(), filepath.Ext(photo.Filename))
		err := SaveFile(photo, userID)
		if err != nil {
			return err
		}
	} else {
		photoData = ""
	}

	data, err := tokoService.TokoRepositoryInterface.FindByIdAndUserID(ctx, id, userID)
	if err != nil {
		return err
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
