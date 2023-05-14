package auth

import (
	"context"
	"errors"
	"time"

	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/exception"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/user"
)

func NewAuthServiceInterface(userRepository *user.UserRepositoryInterface) AuthServiceInterface {
	return &authService{
		UserRepositoryInterface: *userRepository,
	}
}

type authService struct {
	user.UserRepositoryInterface
}

func (authService *authService) Register(ctx context.Context, model model.AuthModel) (string, error) {
	data, err := authService.UserRepositoryInterface.FindByEmail(ctx, model.Email)
	if err != nil {
		return "", err
	}

	if data.Email == model.Email {
		return "", errors.New("Error 1062: Duplicate entry '" + model.Email + "' for key 'users.email'")
	}

	date, err := time.Parse(constants.Layout, model.TanggalLahir)
	exception.PanicLogging(err)

	password, err := common.HashPassword(model.KataSandi)
	if err != nil {
		return "", err
	}

	user := entity.User{
		Nama:         model.Nama,
		KataSandi:    password,
		NoTelp:       model.NoTelp,
		TanggalLahir: date,
		JenisKelamin: model.JenisKelamin,
		Tentang:      model.Tentang,
		Pekerjaan:    model.Pekerjaan,
		Email:        model.Email,
		IdProvinsi:   model.IdProvinsi,
		IdKota:       model.IdKota,
	}

	err = authService.RegisterUser(ctx, user)
	if err != nil {
		return "", err
	}

	return constants.RegisterSuccess, nil
}
