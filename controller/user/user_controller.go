package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/middleware"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/service/user"
)

type UserController struct {
	user.UserServiceInterface
	configuration.Config
}

func NewUserController(userService *user.UserServiceInterface, config configuration.Config) *UserController {
	return &UserController{
		UserServiceInterface: *userService,
		Config:               config,
	}
}

func (controller UserController) Route(app *fiber.App) {
	app.Get("/user", middleware.AuthenticateJWT(controller.Config), controller.GetProfile)
	app.Put("/user", middleware.AuthenticateJWT(controller.Config), controller.UpdateProfile)
}

func (controller UserController) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)

	data, err := controller.UserServiceInterface.GetProfile(c.Context(), userID)
	if err != nil {
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

func (controller UserController) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	var request model.UpdateUserProfileModel
	var errMessage []map[string]interface{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	errMessage = common.Validate(request)
	if len(errMessage) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  errMessage,
			Data:    nil,
		})
	}

	err = controller.UserServiceInterface.UpdateProfile(c.Context(), request, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to PUT data",
		Errors:  nil,
		Data:    "",
	})
}
