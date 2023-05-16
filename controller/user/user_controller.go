package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/middleware"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/service/alamat"
	"github.com/hilmiikhsan/go_rest_api/service/user"
)

type UserController struct {
	user.UserServiceInterface
	configuration.Config
	alamat.AlamatServiceInterface
}

func NewUserController(userService *user.UserServiceInterface, config configuration.Config, alamatService *alamat.AlamatServiceInterface) *UserController {
	return &UserController{
		UserServiceInterface:   *userService,
		Config:                 config,
		AlamatServiceInterface: *alamatService,
	}
}

func (controller UserController) Route(app *fiber.App) {
	app.Get("/user", middleware.AuthenticateJWT(controller.Config), controller.GetProfile)
	app.Put("/user", middleware.AuthenticateJWT(controller.Config), controller.UpdateProfile)
	app.Post("/user/alamat", middleware.AuthenticateJWT(controller.Config), controller.CreateAlamat)
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

func (controller UserController) CreateAlamat(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	var request model.AlamatModel
	var errMessage []map[string]interface{}
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    0,
		})
	}

	errMessage = common.Validate(request)
	if len(errMessage) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  errMessage,
			Data:    0,
		})
	}

	err = controller.AlamatServiceInterface.CreateAlamat(c.Context(), request, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    0,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Errors:  nil,
		Data:    1,
	})
}
