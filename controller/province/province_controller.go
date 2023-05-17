package province

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/service/province"
)

type ProvinceController struct {
	province.ProvinceServiceInterface
	configuration.Config
}

func NewProvinceController(provinceService *province.ProvinceServiceInterface, config configuration.Config) *ProvinceController {
	return &ProvinceController{
		ProvinceServiceInterface: *provinceService,
		Config:                   config,
	}
}

func (controller ProvinceController) Route(app *fiber.App) {
	app.Get("/provcity/listprovincies", controller.GetProvinceCity)
	app.Get("/provcity/detailprovince/:prov_id", controller.GetProvinceDetail)
}

func (controller ProvinceController) GetProvinceCity(c *fiber.Ctx) error {
	data, err := controller.ProvinceServiceInterface.GetProvinceCity(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to get data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to get data",
		Errors:  nil,
		Data:    data,
	})
}

func (controller ProvinceController) GetProvinceDetail(c *fiber.Ctx) error {
	provID := c.Params("prov_id")

	data, err := controller.ProvinceServiceInterface.GetProvinceDetail(c.Context(), provID)
	if err != nil {
		if strings.Contains(err.Error(), "invalid character '<' looking for beginning of value") {
			return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to get data",
				Errors:  []string{"Province not found"},
				Data:    nil,
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to get data",
		Errors:  nil,
		Data:    data,
	})
}
