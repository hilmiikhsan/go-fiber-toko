package auth

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/exception"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/toko"
	"github.com/hilmiikhsan/go_rest_api/repository/user"
	"gorm.io/gorm"
)

func NewAuthServiceInterface(userRepository *user.UserRepositoryInterface, tokoRepository *toko.TokoRepositoryInterface, db *gorm.DB) AuthServiceInterface {
	return &authService{
		UserRepositoryInterface: *userRepository,
		TokoRepositoryInterface: *tokoRepository,
		DB:                      db,
	}
}

type authService struct {
	user.UserRepositoryInterface
	toko.TokoRepositoryInterface
	*gorm.DB
}

func (authService *authService) Register(ctx context.Context, user model.AuthRegisterModel) (string, error) {
	data, err := authService.UserRepositoryInterface.FindByEmail(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if data.Email == user.Email {
		return "", errors.New("Error 1062: Duplicate entry '" + user.Email + "' for key 'users.email'")
	}

	date, err := time.Parse(constants.Layout, user.TanggalLahir)
	exception.PanicLogging(err)

	password, err := common.HashPassword(user.KataSandi)
	if err != nil {
		return "", err
	}

	userData := entity.User{
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

	tx := authService.DB.Begin()

	id, err := authService.UserRepositoryInterface.Insert(ctx, tx, userData)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	tokoData := entity.Toko{
		IdUser:   id,
		NamaToko: user.NamaToko,
	}

	err = authService.TokoRepositoryInterface.Insert(ctx, tx, tokoData)
	if err != nil {
		tx.Rollback()
		return "", err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return "", err
	}

	msgResponse := "Register Succeed"

	return msgResponse, nil
}

func (authService *authService) Login(ctx context.Context, user model.AuthLoginModel) (model.AuthResponseModel, error) {
	response := model.AuthResponseModel{}

	data, err := authService.UserRepositoryInterface.FindByNoTelp(ctx, user.NoTelp)
	if err != nil {
		return model.AuthResponseModel{}, err
	}

	if !common.CheckPasswordHash(user.KataSandi, data.KataSandi) {
		err := constants.ErrPasswordNotMatch
		return model.AuthResponseModel{}, err
	}

	urlProvinsi := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/province/" + data.IdProvinsi + ".json"
	provinceData, err := http.Get(urlProvinsi)
	if err != nil {
		return model.AuthResponseModel{}, err
	}
	defer provinceData.Body.Close()

	var province model.Provinsi
	err = json.NewDecoder(provinceData.Body).Decode(&province)
	if err != nil {
		return model.AuthResponseModel{}, err
	}

	urlKota := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/regency/" + data.IdKota + ".json"
	regencyData, err := http.Get(urlKota)
	if err != nil {
		return model.AuthResponseModel{}, err
	}
	defer regencyData.Body.Close()

	var regency model.Kota
	err = json.NewDecoder(regencyData.Body).Decode(&regency)
	if err != nil {
		return model.AuthResponseModel{}, err
	}

	response = model.AuthResponseModel{
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
