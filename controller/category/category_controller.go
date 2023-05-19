package category

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/middleware"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/service/category"
)

type CategoryController struct {
	category.CategoryServiceInterface
	configuration.Config
}

func NewCategoryController(categoryService *category.CategoryServiceInterface, config configuration.Config) *CategoryController {
	return &CategoryController{
		CategoryServiceInterface: *categoryService,
		Config:                   config,
	}
}

func (controller CategoryController) Route(app *fiber.App) {
	app.Post("/category", middleware.AuthenticateJWT(controller.Config), controller.CreateCategory)
	app.Put("/category/:id", middleware.AuthenticateJWT(controller.Config), controller.UpdateCategoryByID)
}

func (controller CategoryController) CreateCategory(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	var request model.CategoryModel
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

	err = controller.CategoryServiceInterface.CreateCategory(c.Context(), request, userID)
	if err != nil {
		if strings.Contains(err.Error(), "Unauthorized") {
			return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{"Unauthorized"},
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
		Data:    1,
	})
}

func (controller CategoryController) UpdateCategoryByID(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  []string{"id is empty"},
			Data:    nil,
		})
	}

	var request model.CategoryModel
	var errMessage []map[string]interface{}
	err = c.BodyParser(&request)
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

	err = controller.CategoryServiceInterface.UpdateCategoryByID(c.Context(), id, userID, request)
	if err != nil {
		if strings.Contains(err.Error(), "Unauthorized") {
			return c.Status(fiber.StatusUnauthorized).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to PUT data",
				Errors:  []string{"Unauthorized"},
				Data:    nil,
			})
		}

		if strings.Contains(err.Error(), constants.ErrCategoryNotFound.Error()) {
			return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to PUT data",
				Errors:  []string{constants.ErrCategoryNotFound.Error()},
				Data:    nil,
			})
		}

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
