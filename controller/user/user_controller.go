package user

import (
	"github.com/gofiber/fiber/v2"
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
}

func (controller UserController) GetProfile(c *fiber.Ctx) error {
	email := c.Locals("email").(string)

	data, err := controller.UserServiceInterface.GetProfile(c.Context(), email)
	if err != nil {
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
		Data:    data,
	})
}
