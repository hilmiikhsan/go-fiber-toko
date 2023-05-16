package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/exception"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/user"
	"gorm.io/gorm"
)

func NewUserServiceInterface(userRepository *user.UserRepositoryInterface, db *gorm.DB) UserServiceInterface {
	return &userService{
		UserRepositoryInterface: *userRepository,
		DB:                      db,
	}
}

type userService struct {
	user.UserRepositoryInterface
	*gorm.DB
}

func (userService *userService) GetProfile(ctx context.Context, userID int) (model.UserModel, error) {
	response := model.UserModel{}

	data, err := userService.UserRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return response, err
	}

	urlProvinsi := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/province/" + data.IdProvinsi + ".json"
	provinceData, err := http.Get(urlProvinsi)
	if err != nil {
		return response, err
	}
	defer provinceData.Body.Close()

	var province model.Provinsi
	err = json.NewDecoder(provinceData.Body).Decode(&province)
	if err != nil {
		return response, err
	}

	urlKota := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/regency/" + data.IdKota + ".json"
	regencyData, err := http.Get(urlKota)
	if err != nil {
		return response, err
	}
	defer regencyData.Body.Close()

	var regency model.Kota
	err = json.NewDecoder(regencyData.Body).Decode(&regency)
	if err != nil {
		return response, err
	}

	response = model.UserModel{
		ID:           data.ID,
		Nama:         data.Nama,
		NoTelp:       data.NoTelp,
		TanggalLahir: data.TanggalLahir.Format(constants.Layout),
		Tentang:      data.Tentang,
		Pekerjaan:    data.Pekerjaan,
		Email:        data.Email,
		IdProvinsi: model.Provinsi{
			ID:   province.ID,
			Name: province.Name,
		},
		IdKota: model.Kota{
			ID:         regency.ID,
			ProvinceID: regency.ProvinceID,
			Name:       regency.Name,
		},
	}

	return response, nil
}

func (userService *userService) UpdateProfile(ctx context.Context, user model.UpdateUserProfileModel, userID int) error {
	tx := userService.DB.Begin()

	date, err := time.Parse(constants.Layout, user.TanggalLahir)
	exception.PanicLogging(err)

	password, err := common.HashPassword(user.KataSandi)
	if err != nil {
		return err
	}

	userModel := entity.User{
		Nama:         user.Nama,
		KataSandi:    password,
		NoTelp:       user.NoTelp,
		TanggalLahir: date,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
	}

	err = userService.UserRepositoryInterface.Update(ctx, tx, userModel, userID)
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
