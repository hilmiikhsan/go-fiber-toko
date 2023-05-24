package log_product

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewLogProductRepositoryInterface(DB *gorm.DB) LogProductRepositoryInterface {
	return &logProductRepository{
		DB: DB,
	}
}

type logProductRepository struct {
	*gorm.DB
}

func (logProductRepository *logProductRepository) Insert(ctx context.Context, tx *gorm.DB, logProduct entity.LogProduk) (entity.LogProduk, error) {
	err := tx.Create(&logProduct).Error
	if err != nil {
		return entity.LogProduk{}, err
	}

	return logProduct, nil
}

func (logProductRepository *logProductRepository) FindByIdLogProduct(ctx context.Context, idLogProduct int) (entity.LogProduk, error) {
	results := entity.LogProduk{}
	query := logProductRepository.DB.WithContext(ctx).
		Table("log_produk").
		Select("log_produk.*, toko.nama_toko, toko.url_foto, category.name_category").
		Joins("JOIN toko ON log_produk.id_toko = toko.id").
		Joins("JOIN category ON log_produk.id_category = category.id").
		Where("log_produk.id = ?", idLogProduct).
		Order("created_at DESC")

	query = query.Preload("Toko")
	query = query.Preload("Category")

	err := query.Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}

func (logProductRepository *logProductRepository) FindByIDsLogProduct(ctx context.Context, idsLogProduct []int) ([]entity.LogProduk, error) {
	results := []entity.LogProduk{}
	query := logProductRepository.DB.WithContext(ctx).
		Table("log_produk").
		Select("log_produk.*, toko.nama_toko, toko.url_foto, category.name_category").
		Joins("JOIN toko ON log_produk.id_toko = toko.id").
		Joins("JOIN category ON log_produk.id_category = category.id").
		Where("log_produk.id IN ?", idsLogProduct).
		Order("created_at DESC")

	query = query.Preload("Toko")
	query = query.Preload("Category")

	err := query.Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}
