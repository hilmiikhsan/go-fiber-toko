package user

import (
	"context"
	"time"

	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/exception"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/user"
	"github.com/hilmiikhsan/go_rest_api/service/province"
	"gorm.io/gorm"
)

func NewUserServiceInterface(userRepository *user.UserRepositoryInterface, db *gorm.DB, provinceService *province.ProvinceServiceInterface) UserServiceInterface {
	return &userService{
		UserRepositoryInterface:  *userRepository,
		DB:                       db,
		ProvinceServiceInterface: *provinceService,
	}
}

type userService struct {
	user.UserRepositoryInterface
	*gorm.DB
	province.ProvinceServiceInterface
}

func (userService *userService) GetProfile(ctx context.Context, userID int) (model.UserModel, error) {
	response := model.UserModel{}

	data, err := userService.UserRepositoryInterface.FindByID(ctx, userID)
	if err != nil {
		return response, err
	}

	provinceData, err := userService.ProvinceServiceInterface.GetProvinceDetail(ctx, data.IdProvinsi)
	if err != nil {
		return response, err
	}

	cityData, err := userService.ProvinceServiceInterface.GetCityDetail(ctx, data.IdKota)
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
			ID:   provinceData.ID,
			Name: provinceData.Name,
		},
		IdKota: model.Kota{
			ID:         cityData.ID,
			ProvinceID: cityData.ProvinceID,
			Name:       cityData.Name,
		},
	}

	return response, nil
}

func (userService *userService) UpdateProfile(ctx context.Context, user model.UpdateUserProfileModel, userID int) error {
	tx := userService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

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
