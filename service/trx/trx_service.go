package trx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"github.com/hilmiikhsan/go_rest_api/repository/alamat"
	"github.com/hilmiikhsan/go_rest_api/repository/detail_trx"
	"github.com/hilmiikhsan/go_rest_api/repository/foto_produk"
	"github.com/hilmiikhsan/go_rest_api/repository/log_product"
	"github.com/hilmiikhsan/go_rest_api/repository/product"
	"github.com/hilmiikhsan/go_rest_api/repository/trx"
	"gorm.io/gorm"
)

func NewTrxServiceInterface(trxRepository *trx.TrxRepositoryInterface, db *gorm.DB, alamatRepository *alamat.AlamatRepositoryInterface, productRepository *product.ProductRepositoryInterface, logProductRepository *log_product.LogProductRepositoryInterface, detailTrxRepository *detail_trx.DetailTrxRepositoryInterface, fotoProdukRepository *foto_produk.FotoProdukRepositoryInterface) TrxServiceInterface {
	return &trxService{
		TrxRepositoryInterface:        *trxRepository,
		DB:                            db,
		AlamatRepositoryInterface:     *alamatRepository,
		ProductRepositoryInterface:    *productRepository,
		LogProductRepositoryInterface: *logProductRepository,
		DetailTrxRepositoryInterface:  *detailTrxRepository,
		FotoProdukRepositoryInterface: *fotoProdukRepository,
	}
}

type trxService struct {
	trx.TrxRepositoryInterface
	*gorm.DB
	alamat.AlamatRepositoryInterface
	product.ProductRepositoryInterface
	log_product.LogProductRepositoryInterface
	detail_trx.DetailTrxRepositoryInterface
	foto_produk.FotoProdukRepositoryInterface
}

func (trxService *trxService) CreateTrx(ctx context.Context, trx model.TrxModel, userID int) error {
	var hargaTotal, hargaTotalTrx int
	detailTrxModels := []entity.DetailTrx{}
	detailTrxModel := entity.DetailTrx{}
	logProductModel := entity.LogProduk{}

	alamatData, err := trxService.AlamatRepositoryInterface.FindAll(ctx, &struct{ model.ParamsModel }{}, userID)
	if err != nil {
		return err
	}

	found := false
	for _, x := range alamatData {
		if trx.AlamatKirim == x.ID {
			found = true
			break
		}
	}
	if !found {
		return errors.New("Alamat tidak ditemukan")
	}

	tx := trxService.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, detail := range trx.DetailTrx {
		productdata, err := trxService.ProductRepositoryInterface.FindByID(ctx, detail.ProductID)
		if err != nil {
			tx.Rollback()
			return err
		}

		logProductModel = entity.LogProduk{
			IdProduk:      productdata.ID,
			NamaProduk:    productdata.NamaProduk,
			Slug:          productdata.Slug,
			HargaReseller: productdata.HargaReseller,
			HargaKonsumen: productdata.HargaKonsumen,
			Deskripsi:     productdata.Deskripsi,
			IdToko:        productdata.IdToko,
			IdCategory:    productdata.IdCategory,
		}

		logProduct, err := trxService.LogProductRepositoryInterface.Insert(ctx, tx, logProductModel)
		if err != nil {
			tx.Rollback()
			return err
		}

		hargaTotal = calculateHargaTotal(detail.Kuantitas, productdata.HargaKonsumen)
		hargaTotalTrx += hargaTotal

		detailTrxModel = entity.DetailTrx{
			IdTrx:       0,
			IdLogProduk: logProduct.ID,
			IdToko:      logProduct.IdToko,
			Kuantitas:   detail.Kuantitas,
			HargaTotal:  hargaTotal,
		}

		detailTrxModels = append(detailTrxModels, detailTrxModel)
	}

	kodeInvoice := fmt.Sprintf("INV-%d", time.Now().Unix())

	trxModel := entity.Trx{
		IdUser:           userID,
		MethodBayar:      trx.MethodBayar,
		AlamatPengiriman: trx.AlamatKirim,
		HargaTotal:       hargaTotalTrx,
		KodeInvoice:      kodeInvoice,
	}

	idTrx, err := trxService.TrxRepositoryInterface.Insert(ctx, tx, trxModel)
	if err != nil {
		tx.Rollback()
		return err
	}

	for i := range detailTrxModels {
		detailTrxModels[i].IdTrx = idTrx
	}

	err = trxService.DetailTrxRepositoryInterface.BulkInsert(ctx, tx, detailTrxModels)
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

func (trxService *trxService) GetAllTrx(ctx context.Context, params *struct{ model.ParamsTrxModel }, userID int) ([]model.GetTrxModel, error) {
	response := []model.GetTrxModel{}

	if params.Page < 1 {
		params.Page = 1
	}

	if params.Limit < 1 {
		params.Limit = 10
	}

	data, err := trxService.TrxRepositoryInterface.FindAll(ctx, params, userID)
	if err != nil {
		return response, err
	}

	trxIDs := make([]int, len(data))
	for i, x := range data {
		trxIDs[i] = x.ID
	}

	trxDetailData, err := trxService.DetailTrxRepositoryInterface.FindByIDsTrx(ctx, trxIDs)
	if err != nil {
		return response, err
	}

	logProductIDs := make([]int, len(trxDetailData))
	for i, dt := range trxDetailData {
		logProductIDs[i] = dt.IdLogProduk
	}

	logProducts, err := trxService.LogProductRepositoryInterface.FindByIDsLogProduct(ctx, logProductIDs)
	if err != nil {
		return response, err
	}

	logProductMap := make(map[int]model.GetLogProductModel)
	for _, logProduct := range logProducts {
		photosModel := []model.FotoProdukModel{}
		fotoProdukData, err := trxService.FotoProdukRepositoryInterface.FindByProductID(ctx, logProduct.IdProduk)
		if err != nil {
			return response, err
		}

		for _, photo := range fotoProdukData {
			photosModel = append(photosModel, model.FotoProdukModel{
				ID:        photo.ID,
				ProductID: photo.IdProduk,
				Url:       photo.Url,
			})
		}

		logProductMap[logProduct.ID] = model.GetLogProductModel{
			ID:            logProduct.IdProduk,
			NamaProduk:    logProduct.NamaProduk,
			Slug:          logProduct.Slug,
			HargaReseller: logProduct.HargaReseller,
			HargaKonsumen: logProduct.HargaKonsumen,
			Deskripsi:     logProduct.Deskripsi,
			Toko: model.GetTokoDetailModel{
				NamaToko: logProduct.Toko.NamaToko,
				UrlFoto:  logProduct.Toko.UrlFoto,
			},
			Category: model.GetCategoryModel{
				ID:           logProduct.Category.ID,
				NamaCategory: logProduct.Category.NamaCategory,
			},
			Photos: photosModel,
		}
	}

	for _, x := range data {
		detailTrxDatas := []model.GetDetailTrxModel{}

		for _, dt := range trxDetailData {
			logProduct := logProductMap[dt.IdLogProduk]

			detailTrxDatas = append(detailTrxDatas, model.GetDetailTrxModel{
				Product: logProduct,
				Toko: model.GetTokoModel{
					ID:       dt.IdToko,
					NamaToko: logProduct.Toko.NamaToko,
					UrlFoto:  logProduct.Toko.UrlFoto,
				},
				Kuantitas:  dt.Kuantitas,
				HargaTotal: dt.HargaTotal,
			})
		}

		response = append(response, model.GetTrxModel{
			ID:          x.ID,
			HargaTotal:  x.HargaTotal,
			KodeInvoice: x.KodeInvoice,
			MethodBayar: x.MethodBayar,
			AlamatKirim: model.GetAlamatModel{
				ID:           x.Alamat.ID,
				JudulAlamat:  x.Alamat.JudulAlamat,
				NamaPenerima: x.Alamat.NamaPenerima,
				NoTelp:       x.Alamat.NoTelp,
				DetailAlamat: x.Alamat.DetailAlamat,
			},
			DetailTrx: detailTrxDatas,
		})
	}

	return response, nil
}

func calculateHargaTotal(kuantitas, hargaKonsumen int) int {
	return kuantitas * hargaKonsumen
}
