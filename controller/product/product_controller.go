package product

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
	"github.com/hilmiikhsan/go_rest_api/service/product"
)

type ProductController struct {
	product.ProductServiceInterface
	configuration.Config
}

func NewProductController(producService *product.ProductServiceInterface, config configuration.Config) *ProductController {
	return &ProductController{
		ProductServiceInterface: *producService,
		Config:                  config,
	}
}

func (controller ProductController) Route(app *fiber.App) {
	app.Post("/product", middleware.AuthenticateJWT(controller.Config), controller.CreateProduct)
}

func (controller ProductController) CreateProduct(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)
	var photos []*multipart.FileHeader
	var request model.CreateProductModel
	namaProduk := c.FormValue("nama_produk")
	categoryIdStr := c.FormValue("category_id")
	hargaResellerStr := c.FormValue("harga_reseller")
	hargaKonsumenStr := c.FormValue("harga_konsumen")
	stokStr := c.FormValue("stok")
	deskripsi := c.FormValue("deskripsi")

	categoryID, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  []string{"category id is empty"},
			Data:    nil,
		})
	}

	hargaReseller, err := strconv.Atoi(hargaResellerStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  []string{"harga reseller is empty"},
			Data:    nil,
		})
	}

	hargaKonsumen, err := strconv.Atoi(hargaKonsumenStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  []string{"harga konsumen is empty"},
			Data:    nil,
		})
	}

	stok, err := strconv.Atoi(stokStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to PUT data",
			Errors:  []string{"stok is empty"},
			Data:    nil,
		})
	}

	request = model.CreateProductModel{
		NamaProduk:    namaProduk,
		CategoryID:    categoryID,
		HargaReseller: hargaReseller,
		HargaKonsumen: hargaKonsumen,
		Stok:          stok,
		Deskripsi:     deskripsi,
	}

	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{"photos is empty"},
			Data:    nil,
		})
	}

	photos = form.File["photos"]

	allowedExtensions := []string{".jpg", ".jpeg", ".png"}
	invalidExtensions := make([]string, 0)

	for _, photo := range photos {
		validExtension := false

		for _, ext := range allowedExtensions {
			if strings.ToLower(filepath.Ext(photo.Filename)) == ext {
				validExtension = true
				break
			}
		}

		if !validExtension {
			invalidExtensions = append(invalidExtensions, photo.Filename)
		}
	}

	if len(invalidExtensions) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
			Status:  false,
			Message: "Failed to POST data",
			Errors:  []string{"Invalid file extension"},
			Data:    nil,
		})
	}

	maxFileSize := 1 * 1024 * 1024
	for _, pt := range photos {
		if pt.Size > int64(maxFileSize) {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{"Photo size exceeds the limit"},
				Data:    nil,
			})
		}
	}

	err = controller.ProductServiceInterface.CreateProduct(c.Context(), request, photos, userID)
	if err != nil {
		if strings.Contains(err.Error(), constants.ErrNamaProdukIsRequired.Error()) {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{"nama produk is empty"},
				Data:    nil,
			})
		}

		if strings.Contains(err.Error(), constants.ErrDeskripsiIsRequired.Error()) {
			return c.Status(fiber.StatusBadRequest).JSON(model.GeneralResponse{
				Status:  false,
				Message: "Failed to POST data",
				Errors:  []string{"deskripsi is empty"},
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
		Data:    4,
	})
}
