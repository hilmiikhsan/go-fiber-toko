package trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"gorm.io/gorm"
)

func NewProductRepositoryInterface(DB *gorm.DB) TrxRepositoryInterface {
	return &trxRepository{
		DB: DB,
	}
}

type trxRepository struct {
	*gorm.DB
}

func (trxRepository *trxRepository) Insert(ctx context.Context, tx *gorm.DB, trx entity.Trx) (int, error) {
	err := tx.WithContext(ctx).Create(&trx).Error
	if err != nil {
		return 0, err
	}

	return trx.ID, nil
}

func (trxRepository *trxRepository) FindAll(ctx context.Context, params *struct{ model.ParamsTrxModel }, userID int) ([]entity.Trx, error) {
	results := []entity.Trx{}
	query := trxRepository.DB.WithContext(ctx).
		Table("trx").
		Select("trx.*, alamat.id as alamat_id, alamat.judul_alamat, alamat.nama_penerima, alamat.no_telp, alamat.detail_alamat").
		Joins("JOIN alamat ON trx.alamat_pengiriman = alamat.id").
		Order("created_at DESC").
		Where("trx.id_user = ?", userID)

	query = query.Preload("Alamat")

	var totalRows int64
	offset := (params.Page - 1) * params.Limit

	if params.Search != "" {
		query = query.Where("kode_invoice LIKE ?", "%"+params.Search+"%").Or("method_bayar LIKE ?", "%"+params.Search+"%")
	}

	err := query.Model(&entity.Produk{}).Count(&totalRows).Error
	if err != nil {
		return results, err
	}

	err = query.Offset(offset).Limit(params.Limit).Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}
