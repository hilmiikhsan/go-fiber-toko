package province

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hilmiikhsan/go_rest_api/model"
)

func NewProvinceServiceInterface() ProvinceServiceInterface {
	return &provinceService{}
}

type provinceService struct {
}

func (provinceService *provinceService) GetProvinceCity(ctx context.Context) ([]model.Provinsi, error) {
	response := []model.Provinsi{}

	urlProvinsi := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/provinces.json"
	provinceData, err := http.Get(urlProvinsi)
	if err != nil {
		return response, err
	}
	defer provinceData.Body.Close()

	var provinces []model.Provinsi
	err = json.NewDecoder(provinceData.Body).Decode(&provinces)
	if err != nil {
		return response, err
	}

	for _, x := range provinces {
		response = append(response, model.Provinsi{
			ID:   x.ID,
			Name: x.Name,
		})
	}

	return response, nil
}

func (provinceService *provinceService) GetProvinceDetail(ctx context.Context, provID string) (model.Provinsi, error) {
	response := model.Provinsi{}

	urlProvinsi := "https://hilmiikhsan.github.io/api-wilayah-indonesia/api/province/" + provID + ".json"
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

	response = model.Provinsi{
		ID:   province.ID,
		Name: province.Name,
	}

	return response, nil
}
