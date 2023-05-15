package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/user"
)

func NewUserServiceInterface(userRepository *user.UserRepositoryInterface) UserServiceInterface {
	return &userService{
		UserRepositoryInterface: *userRepository,
	}
}

type userService struct {
	user.UserRepositoryInterface
}

func (userService *userService) GetProfile(ctx context.Context, email string) (model.User, error) {
	response := model.User{}

	data, err := userService.UserRepositoryInterface.FindByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	urlProvinsi := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/province/" + data.IdProvinsi + ".json"
	provinceData, err := http.Get(urlProvinsi)
	if err != nil {
		return model.User{}, err
	}
	defer provinceData.Body.Close()

	var province model.Provinsi
	err = json.NewDecoder(provinceData.Body).Decode(&province)
	if err != nil {
		return model.User{}, err
	}

	urlKota := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/regency/" + data.IdKota + ".json"
	regencyData, err := http.Get(urlKota)
	if err != nil {
		return model.User{}, err
	}
	defer regencyData.Body.Close()

	var regency model.Kota
	err = json.NewDecoder(regencyData.Body).Decode(&regency)
	if err != nil {
		return model.User{}, err
	}

	response = model.User{
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
