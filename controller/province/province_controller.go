package province

import (
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
