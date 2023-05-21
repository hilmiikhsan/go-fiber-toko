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
	"github.com/hilmiikhsan/go_rest_api/repository/foto_produk"
	"github.com/hilmiikhsan/go_rest_api/repository/product"
	"github.com/hilmiikhsan/go_rest_api/repository/toko"
	"gorm.io/gorm"
)

func NewProductServiceInterface(productRepository *product.ProductRepositoryInterface, db *gorm.DB, tokoRepository *toko.TokoRepositoryInterface, fotoProdukRepository *foto_produk.FotoProdukRepositoryInterface) ProductServiceInterface {
	return &productService{
		ProductRepositoryInterface:    *productRepository,
		DB:                            db,
		TokoRepositoryInterface:       *tokoRepository,
		FotoProdukRepositoryInterface: *fotoProdukRepository,
	}
}

type productService struct {
	product.ProductRepositoryInterface
	*gorm.DB
	toko.TokoRepositoryInterface
	foto_produk.FotoProdukRepositoryInterface
}

func (productService *productService) CreateProduct(ctx context.Context, product model.CreateProductModel, photos []*multipart.FileHeader, userID int) error {
	var photoData string
	var fotoProdukModel entity.FotoProduk

	if product.NamaProduk == "" {
		return errors.New("Nama Produk is required")
	}

	if product.Deskripsi == "" {
		return errors.New("Deskripsi is required")
	}

	tokoData, err := productService.TokoRepositoryInterface.FindByUserID(ctx, userID)
	if err != nil {
		return err
	}

	tx := productService.DB.Begin()

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
		return err
	}

	return nil
}
