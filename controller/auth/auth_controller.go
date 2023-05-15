package auth

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/service/auth"
)

func NewAuthController(authService *auth.AuthServiceInterface, config configuration.Config) *AuthController {
	return &AuthController{
		AuthServiceInterface: *authService,
		Config:               config,
	}
}

type AuthController struct {
	auth.AuthServiceInterface
	configuration.Config
}

func (controller AuthController) Route(app *fiber.App) {
	app.Post("/auth/register", controller.Register)
	app.Post("/auth/login", controller.Login)
}

func (controller AuthController) Register(c *fiber.Ctx) error {
	var request model.AuthRegisterModel
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	response, err := controller.AuthServiceInterface.Register(c.Context(), request)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062: Duplicate entry '"+request.Email+"' for key 'users.email'") {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{"Error 1062: Duplicate entry '" + request.Email + "' for key 'users.email'"},
				Data:    nil,
			})
		}

		if err, ok := err.(*mysql.MySQLError); ok {
			if err.Number == 1062 {
				errMsg := fmt.Sprintf("Duplicate entry for %s", err.Message)
				return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
					Status:  false,
					Message: "Failed to POST data",
					Errors:  []string{errMsg},
					Data:    nil,
				})
			}
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
		Data:    response,
	})
}

func (controller AuthController) Login(c *fiber.Ctx) error {
	var request model.AuthLoginModel
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	response, err := controller.AuthServiceInterface.Login(c.Context(), request)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrUserNotFound.Error()) || strings.Contains(err.Error(), constants.ErrPasswordNotMatch.Error()) {
			return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{"No Telp atau kata sandi salah"},
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

	token := common.GenerateToken(response.NoTelp, controller.Config)
	resultWithToken := map[string]interface{}{
		"token": token,
	}

	response.Token = resultWithToken

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to POST data",
		Errors:  nil,
		Data:    response,
	})
}
