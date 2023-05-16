package alamat

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"github.com/hilmiikhsan/go_rest_api/model"
	"gorm.io/gorm"
)

func NewAlamatRepositoryInterface(DB *gorm.DB) AlamatRepositoryInterface {
	return &alamatRepository{
		DB: DB,
	}
}

type alamatRepository struct {
	*gorm.DB
}

func (alamatRepository *alamatRepository) Insert(ctx context.Context, tx *gorm.DB, alamat entity.Alamat) error {
	err := tx.WithContext(ctx).Create(&alamat).Error
	if err != nil {
		return err
	}

	return nil
}

func (alamatRepository *alamatRepository) FindAll(ctx context.Context, params *struct{ model.ParamsModel }, userID int) ([]entity.Alamat, error) {
	results := []entity.Alamat{}
	query := alamatRepository.DB
	if params.JudulAlamat != "" {
		query = query.Where("alamat.judul_alamat LIKE ?", "%"+params.JudulAlamat+"%").Where("alamat.id_user = ?", userID)
	}

	err := query.Where("alamat.id_user", userID).Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}

func (alamatRepository *alamatRepository) FindByID(ctx context.Context, id, userID int) (entity.Alamat, error) {
	alamat := entity.Alamat{}
	result := alamatRepository.DB.WithContext(ctx).Where("alamat.id = ?", id).Where("alamat.id_user = ?", userID).Find(&alamat)
	if result.RowsAffected == 0 {
		return alamat, errors.New("record not found")
	}

	return alamat, nil
}

func (alamatRepository *alamatRepository) Update(ctx context.Context, tx *gorm.DB, alamat entity.Alamat, id, userID int) error {
	err := tx.WithContext(ctx).Where("alamat.id = ?", id).Where("alamat.id_user = ?", userID).Updates(&alamat).Error
	if err != nil {
		return err
	}

	return nil
}

func (alamatRepository *alamatRepository) Delete(ctx context.Context, tx *gorm.DB, alamat entity.Alamat, id, userID int) error {
	err := tx.WithContext(ctx).Where("alamat.id = ?", id).Where("alamat.id_user = ?", userID).Delete(&alamat).Error
	if err != nil {
		return err
	}

	return nil
}
