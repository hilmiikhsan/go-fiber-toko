package detail_trx

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewDetailTrxRepository(DB *gorm.DB) DetailTrxRepositoryInterface {
	return &detailTrxRepository{
		DB: DB,
	}
}

type detailTrxRepository struct {
	*gorm.DB
}

func (detailTrxRepository *detailTrxRepository) BulkInsert(ctx context.Context, tx *gorm.DB, detailTrx []entity.DetailTrx) error {
	err := tx.Create(&detailTrx).Error
	if err != nil {
		return err
	}

	return nil
}

func (detailTrxRepository *detailTrxRepository) FindByIdTrx(ctx context.Context, idTrx int) ([]entity.DetailTrx, error) {
	results := []entity.DetailTrx{}
	err := detailTrxRepository.DB.WithContext(ctx).Where("detail_trx.id_trx = ?", idTrx).Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}

func (detailTrxRepository *detailTrxRepository) FindByIDsTrx(ctx context.Context, idsTrx []int) ([]entity.DetailTrx, error) {
	results := []entity.DetailTrx{}
	err := detailTrxRepository.DB.WithContext(ctx).Where("detail_trx.id_trx IN ?", idsTrx).Find(&results).Error
	if err != nil {
		return results, err
	}

	return results, nil
}
