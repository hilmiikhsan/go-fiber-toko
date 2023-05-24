package product

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/gosimple/slug"
	"github.com/hilmiikhsan/go_rest_api/common"
	"github.com/hilmiikhsan/go_rest_api/constants"
	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/category"
	"github.com/hilmiikhsan/go_rest_api/repository/foto_produk"
	"github.com/hilmiikhsan/go_rest_api/repository/product"
	"github.com/hilmiikhsan/go_rest_api/repository/toko"
	"gorm.io/gorm"
)

func NewProductServiceInterface(productRepository *product.ProductRepositoryInterface, db *gorm.DB, tokoRepository *toko.TokoRepositoryInterface, fotoProdukRepository *foto_produk.FotoProdukRepositoryInterface, categoryRepository *category.CategoryRepositoryInterface) ProductServiceInterface {
	return &productService{
		ProductRepositoryInterface:    *productRepository,
		DB:                            db,
		TokoRepositoryInterface:       *tokoRepository,
		FotoProdukRepositoryInterface: *fotoProdukRepository,
		CategoryRepositoryInterface:   *categoryRepository,
	}
}

type productService struct {
	product.ProductRepositoryInterface
	*gorm.DB
	toko.TokoRepositoryInterface
	foto_produk.FotoProdukRepositoryInterface
	category.CategoryRepositoryInterface
}

func (productService *productService) CreateProduct(ctx context.Context, product model.ProductModel, photos []*multipart.FileHeader, userID int) error {
	var photoData string
	var fotoProdukModel entity.FotoProduk

	if product.NamaProduk == "" {
		return errors.New("Nama Produk is required")
	}

	if product.Deskripsi == "" {
		return errors.New("Deskripsi is required")
	}

	_, err := productService.CategoryRepositoryInterface.FindByID(ctx, product.CategoryID)
	if err != nil {
		return err
	}

	tokoData, err := productService.TokoRepositoryInterface.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}

	tx := productService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	productSlug := slug.Make(product.NamaProduk)

	productModel := entity.Produk{
		NamaProduk:    product.NamaProduk,
		Slug:          productSlug,
		IdCategory:    product.CategoryID,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		IdToko:        tokoData.ID,
	}

	productID, err := productService.ProductRepositoryInterface.Insert(ctx, tx, productModel)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, photo := range photos {
		if len(photo.Filename) > 0 {
			photoData = fmt.Sprintf("%d-%s", common.GenerateUniqueID(), photo.Filename)
			err := common.SaveFile(photo, constants.TemporaryProductFilePath)
			if err != nil {
				tx.Rollback()
				return err
			}

			fotoProdukModel = entity.FotoProduk{
				IdProduk: productID,
				Url:      photoData,
			}

			err = productService.FotoProdukRepositoryInterface.Insert(ctx, tx, fotoProdukModel)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (productService *productService) UpdateProductByID(ctx context.Context, product model.ProductModel, photos []*multipart.FileHeader, id, userID int) error {
	var photoData string
	var fotoProdukModel entity.FotoProduk

	if product.NamaProduk == "" {
		return errors.New("Nama Produk is required")
	}

	if product.Deskripsi == "" {
		return errors.New("Deskripsi is required")
	}

	productData, err := productService.ProductRepositoryInterface.FindByID(ctx, id)
	if err != nil {
		return err
	}

	tokoData, err := productService.TokoRepositoryInterface.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}

	tx := productService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	productSlug := slug.Make(product.NamaProduk)

	productModel := entity.Produk{
		NamaProduk:    product.NamaProduk,
		Slug:          productSlug,
		IdCategory:    product.CategoryID,
		HargaReseller: product.HargaReseller,
		HargaKonsumen: product.HargaKonsumen,
		Stok:          product.Stok,
		Deskripsi:     product.Deskripsi,
		IdToko:        tokoData.ID,
	}

	err = productService.ProductRepositoryInterface.Update(ctx, tx, productModel, productData.ID, tokoData.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	fotoProdukData, err := productService.FotoProdukRepositoryInterface.FindByProductID(ctx, productData.ID)
	if err != nil {
		return err
	}

	for i, photo := range photos {
		if len(photo.Filename) > 0 {
			photoData = fmt.Sprintf("%d-%s", common.GenerateUniqueID(), photo.Filename)
			err := common.SaveFile(photo, constants.TemporaryProductFilePath)
			if err != nil {
				tx.Rollback()
				return err
			}

			fotoProdukModel = entity.FotoProduk{
				Url: photoData,
			}

			err = productService.FotoProdukRepositoryInterface.Update(ctx, tx, fotoProdukModel, fotoProdukData[i].ID, productData.ID)
			if err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (productService *productService) DeleteProductByID(ctx context.Context, id, userID int) error {
	productData, err := productService.ProductRepositoryInterface.FindByID(ctx, id)
	if err != nil {
		return err
	}

	tokoData, err := productService.TokoRepositoryInterface.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}

	tx := productService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err = productService.ProductRepositoryInterface.Delete(ctx, tx, productData, productData.ID, tokoData.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (productService *productService) GetAllProduct(ctx context.Context, params *struct{ model.ParamsProductModel }) ([]model.GetProductModel, error) {
	response := []model.GetProductModel{}
	tokoIDs := []int{}

	if params.Page < 1 {
		params.Page = 1
	}

	if params.Limit < 1 {
		params.Limit = 10
	}

	data, err := productService.ProductRepositoryInterface.FindAll(ctx, params)
	if err != nil {
		return response, err
	}

	for _, dt := range data {
		tokoIDs = append(tokoIDs, dt.IdToko)
	}

	photos, err := productService.FotoProdukRepositoryInterface.FindAll(ctx, tokoIDs)
	if err != nil {
		return response, err
	}

	photosByProductID := make(map[int][]entity.FotoProduk)
	for _, photo := range photos {
		photosByProductID[photo.Produk.ID] = append(photosByProductID[photo.Produk.ID], photo)
	}

	for _, x := range data {
		photosModel := []model.FotoProdukModel{}
		for _, photo := range photosByProductID[x.ID] {
			photosModel = append(photosModel, model.FotoProdukModel{
				ID:        photo.ID,
				ProductID: photo.IdProduk,
				Url:       photo.Url,
			})
		}

		response = append(response, model.GetProductModel{
			ID:            x.ID,
			NamaProduk:    x.NamaProduk,
			Slug:          x.Slug,
			HargaReseller: x.HargaReseller,
			HargaKonsumen: x.HargaKonsumen,
			Stok:          x.Stok,
			Deskripsi:     x.Deskripsi,
			Toko: model.GetTokoModel{
				ID:       x.IdToko,
				NamaToko: x.Toko.NamaToko,
				UrlFoto:  x.Toko.UrlFoto,
			},
			Category: model.GetCategoryModel{
				ID:           x.IdCategory,
				NamaCategory: x.Category.NamaCategory,
			},
			Photos: photosModel,
		})
	}

	return response, nil
}

func (productService *productService) GetProductByID(ctx context.Context, id int) (model.GetProductModel, error) {
	response := model.GetProductModel{}
	photoDatas := []model.FotoProdukModel{}

	data, err := productService.ProductRepositoryInterface.FindByID(ctx, id)
	if err != nil {
		return response, err
	}

	photos, err := productService.FotoProdukRepositoryInterface.FindByProductID(ctx, data.ID)
	if err != nil {
		return response, err
	}

	for _, photo := range photos {
		photoDatas = append(photoDatas, model.FotoProdukModel{
			ID:        photo.ID,
			ProductID: photo.IdProduk,
			Url:       photo.Url,
		})
	}

	response = model.GetProductModel{
		ID:            data.ID,
		NamaProduk:    data.NamaProduk,
		Slug:          data.Slug,
		HargaReseller: data.HargaReseller,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		Toko: model.GetTokoModel{
			ID:       data.Toko.ID,
			NamaToko: data.Toko.NamaToko,
		},
		Category: model.GetCategoryModel{
			ID:           data.IdToko,
			NamaCategory: data.Category.NamaCategory,
		},
		Photos: photoDatas,
	}

	return response, nil
}
