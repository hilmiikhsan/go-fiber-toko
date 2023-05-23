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
	"github.com/hilmiikhsan/go_rest_api/repository/log_product"
	"github.com/hilmiikhsan/go_rest_api/repository/product"
	"github.com/hilmiikhsan/go_rest_api/repository/trx"
	"gorm.io/gorm"
)

func NewTrxServiceInterface(trxRepository *trx.TrxRepositoryInterface, db *gorm.DB, alamatRepository *alamat.AlamatRepositoryInterface, productRepository *product.ProductRepositoryInterface, logProductRepository *log_product.LogProductRepositoryInterface, detailTrxRepository *detail_trx.DetailTrxRepositoryInterface) TrxServiceInterface {
	return &trxService{
		TrxRepositoryInterface:        *trxRepository,
		DB:                            db,
		AlamatRepositoryInterface:     *alamatRepository,
		ProductRepositoryInterface:    *productRepository,
		LogProductRepositoryInterface: *logProductRepository,
		DetailTrxRepositoryInterface:  *detailTrxRepository,
	}
}

type trxService struct {
	trx.TrxRepositoryInterface
	*gorm.DB
	alamat.AlamatRepositoryInterface
	product.ProductRepositoryInterface
	log_product.LogProductRepositoryInterface
	detail_trx.DetailTrxRepositoryInterface
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

func calculateHargaTotal(kuantitas, hargaKonsumen int) int {
	return kuantitas * hargaKonsumen
}
