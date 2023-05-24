package product

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"gorm.io/gorm"
)

func NewProductRepositoryInterface(DB *gorm.DB) ProductRepositoryInterface {
	return &productRepository{
		DB: DB,
	}
}

type productRepository struct {
	*gorm.DB
}

func (productRepository *productRepository) Insert(ctx context.Context, tx *gorm.DB, product entity.Produk) (int, error) {
	err := tx.WithContext(ctx).Create(&product).Error
	if err != nil {
		return 0, err
	}

	return product.ID, nil
}

func (productRepository *productRepository) Update(ctx context.Context, tx *gorm.DB, product entity.Produk, id, idToko int) error {
	err := tx.WithContext(ctx).Where("produk.id = ?", id).Where("produk.id_toko = ?", idToko).Updates(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (productRepository *productRepository) FindByID(ctx context.Context, id int) (entity.Produk, error) {
	product := entity.Produk{}
	query := productRepository.DB.WithContext(ctx).
		Table("produk").
		Select("produk.*, toko.nama_toko, toko.url_foto, category.name_category").
		Joins("JOIN toko ON produk.id_toko = toko.id").
		Joins("JOIN category ON produk.id_category = category.id").
		Where("produk.id = ?", id)

	query = query.Preload("Toko")
	query = query.Preload("Category")

	result := query.Find(&product)
	if result.RowsAffected == 0 {
		return entity.Produk{}, errors.New("No Data Product")
	}

	return product, nil
}

func (productRepository *productRepository) Delete(ctx context.Context, tx *gorm.DB, product entity.Produk, id, idToko int) error {
	err := tx.WithContext(ctx).Where("produk.id = ?", id).Where("produk.id_toko = ?", idToko).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (productRepository *productRepository) FindAll(ctx context.Context, params *struct{ model.ParamsProductModel }) ([]entity.Produk, error) {
	results := []entity.Produk{}
	query := productRepository.DB.WithContext(ctx).
		Table("produk").
		Select("produk.*, toko.nama_toko, toko.url_foto, category.name_category").
		Joins("JOIN toko ON produk.id_toko = toko.id").
		Joins("JOIN category ON produk.id_category = category.id").
		Order("created_at DESC")

	query = query.Preload("Toko")
	query = query.Preload("Category")

	var totalRows int64
	offset := (params.Page - 1) * params.Limit

	if params.NamaProduk != "" {
		query = query.Where("produk.nama_produk LIKE ?", "%"+params.NamaProduk+"%")
	}

	if params.CategoryID > 0 {
		query = query.Where("produk.id_category = ?", params.CategoryID)
	}

	if params.TokoID > 0 {
		query = query.Where("produk.id_toko = ?", params.TokoID)
	}

	if params.MaxHarga > 0 {
		query = query.Where("produk.harga_konsumen <= ?", params.MaxHarga)
	}

	if params.MinHarga > 0 {
		query = query.Where("produk.harga_konsumen >= ?", params.MinHarga)
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

func (productRepository *productRepository) FindByIDs(ctx context.Context, IDs []int) ([]entity.Produk, error) {
	results := []entity.Produk{}
	err := productRepository.DB.WithContext(ctx).Where("produk.id IN ?", IDs).Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}
