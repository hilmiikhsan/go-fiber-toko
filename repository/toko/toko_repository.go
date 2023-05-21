package toko

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"gorm.io/gorm"
)

func NewTokoRepositoryInterface(DB *gorm.DB) TokoRepositoryInterface {
	return &tokoRepository{
		DB: DB,
	}
}

type tokoRepository struct {
	*gorm.DB
}

func (tokoRepository *tokoRepository) Insert(ctx context.Context, tx *gorm.DB, toko entity.Toko) error {
	err := tx.WithContext(ctx).Create(&toko).Error
	if err != nil {
		return err
	}

	return nil
}

func (tokoRepository *tokoRepository) FindByUserID(ctx context.Context, id int) (entity.Toko, error) {
	toko := entity.Toko{}
	result := tokoRepository.DB.WithContext(ctx).Where("toko.id_user = ?", id).Find(&toko)
	if result.RowsAffected == 0 {
		return entity.Toko{}, errors.New("record not found")
	}

	return toko, nil
}

func (tokoRepository *tokoRepository) Update(ctx context.Context, tx *gorm.DB, toko entity.Toko, id, userID int) error {
	err := tx.WithContext(ctx).Where("toko.id = ?", id).Where("toko.id_user = ?", userID).Updates(&toko).Error
	if err != nil {
		return err
	}

	return nil
}

func (tokoRepository *tokoRepository) FindByIdAndUserID(ctx context.Context, id, userID int) (entity.Toko, error) {
	toko := entity.Toko{}
	result := tokoRepository.DB.WithContext(ctx).Where("toko.id = ?", id).Where("toko.id_user = ?", userID).Find(&toko)
	if result.RowsAffected == 0 {
		return entity.Toko{}, errors.New("record not found")
	}

	return toko, nil
}

func (tokoRepository *tokoRepository) FindAll(ctx context.Context, params *struct{ model.ParamsTokoModel }) ([]entity.Toko, error) {
	results := []entity.Toko{}
	query := tokoRepository.DB.Joins("JOIN user ON toko.id_user = user.id").Where("user.isAdmin = ?", 0)
	var totalRows int64
	offset := (params.Page - 1) * params.Limit

	if params.Nama != "" {
		query = query.Where("toko.nama_toko LIKE ?", "%"+params.Nama+"%")
	}

	err := query.Model(&entity.Toko{}).Count(&totalRows).Error
	if err != nil {
		return results, err
	}

	err = query.Offset(offset).Limit(params.Limit).Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}

func (tokoRepository *tokoRepository) FindByID(ctx context.Context, id int) (entity.Toko, error) {
	toko := entity.Toko{}
	result := tokoRepository.DB.WithContext(ctx).Where("id = ?", id).Find(&toko)
	if result.RowsAffected == 0 {
		return entity.Toko{}, errors.New("Toko tidak ditemukan")
	}

	return toko, nil
}
