package toko

import (
	"mime/multipart"
	"path/filepath"
	"strconv"
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
	app.Put("/toko/:id_toko", middleware.AuthenticateJWT(controller.Config), controller.UpdateToko)
	app.Get("/toko", middleware.AuthenticateJWT(controller.Config), controller.GetAllToko)
	app.Get("/toko/:id_toko", middleware.AuthenticateJWT(controller.Config), controller.GetTokoByID)
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

func (controller TokoController) UpdateToko(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	idStr := c.Params("id_toko")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to UPDATE data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	var photo *multipart.FileHeader

	namaToko := c.FormValue("nama_toko")
	photo, err = c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to UPDATE data",
			Errors:  []string{"Photo is empty"},
			Data:    nil,
		})
	}

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	validExtension := false
	for _, ext := range allowedExtensions {
		if strings.ToLower(filepath.Ext(photo.Filename)) == ext {
			validExtension = true
			break
		}
	}

	if !validExtension {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to UPDATE data",
			Errors:  []string{"Invalid file extension"},
			Data:    nil,
		})
	}

	maxFileSize := 1 * 1024 * 1024
	if photo.Size > int64(maxFileSize) {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to UPDATE data",
			Errors:  []string{"Photo size exceeds the limit"},
			Data:    nil,
		})
	}

	err = controller.TokoServiceInterface.UpdateToko(c.Context(), namaToko, photo, id, userID)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrNamaTokoIsRequired.Error()) {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to UPDATE data",
				Errors:  []string{err.Error()},
				Data:    nil,
			})
		}

		if strings.Contains(err.Error(), constants.ErrRecordNotFound.Error()) {
			return c.Status(fiber.StatusNotFound).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to UPDATE data",
				Errors:  []string{err.Error()},
				Data:    nil,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to UPDATE data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(model.GeneralResponse{
		Status:  true,
		Message: "Succeed to UPDATE data",
		Errors:  nil,
		Data:    "Update toko succeed",
	})
}

func (controller TokoController) GetAllToko(c *fiber.Ctx) error {
	params := new(struct {
		model.ParamsTokoModel
	})

	err := c.QueryParser(params)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to GET data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	data, err := controller.TokoServiceInterface.GetAllToko(c.Context(), params)
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
		Data: model.PaginationResponse{
			Page:  params.Page,
			Limit: params.Limit,
			Data:  data,
		},
	})
}

func (controller TokoController) GetTokoByID(c *fiber.Ctx) error {
	idStr := c.Params("id_toko")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to GET data",
			Errors:  []string{err.Error()},
			Data:    nil,
		})
	}

	data, err := controller.TokoServiceInterface.GeTokoByID(c.Context(), id)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrTokoNotFound.Error()) {
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
