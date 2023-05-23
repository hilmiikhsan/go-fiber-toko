package trx

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/middleware"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/service/trx"
)

type TrxController struct {
	trx.TrxServiceInterface
	configuration.Config
}

func NewTrxController(trxService *trx.TrxServiceInterface, config configuration.Config) *TrxController {
	return &TrxController{
		TrxServiceInterface: *trxService,
		Config:              config,
	}
}

func (controller TrxController) Route(app *fiber.App) {
	app.Post("/trx", middleware.AuthenticateJWT(controller.Config), controller.CreateTrx)
}

func (controller TrxController) CreateTrx(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	var request model.TrxModel
	var errMessage []map[string]interface{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	errMessage = common.Validate(request)
	if len(errMessage) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  errMessage,
			Data:    nil,
		})
	}

	err = controller.TrxServiceInterface.CreateTrx(c.Context(), request, userID)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrAlamatNotFound.Error()) {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{err.Error()},
				Data:    nil,
			})
		}

		if strings.Contains(err.Error(), constants.ErrProductNotFound.Error()) {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{err.Error()},
				Data:    nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Errors:  nil,
		Data:    6,
	})
}