package toko

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/middleware"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/service/toko"
)

type TokoController struct {
	toko.TokoServiceInterface
	configuration.Config
}

func NewTokoController(tokoService *toko.TokoServiceInterface, config configuration.Config) *TokoController {
	return &TokoController{
		TokoServiceInterface: *tokoService,
		Config:               config,
	}
}

func (controller TokoController) Route(app *fiber.App) {
	app.Get("/toko/my", middleware.AuthenticateJWT(controller.Config), controller.GetMyToko)
}

func (controller TokoController) GetMyToko(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)

	data, err := controller.TokoServiceInterface.GetMyToko(c.Context(), userID)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrRecordNotFound.Error()) {
			return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to GET data",
				Errors:  []string{err.Error()},
				Data:    nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to GET data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to GET data",
		Errors:  nil,
		Data:    data,
	})
}
