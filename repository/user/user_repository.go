package user

import (
	"context"
	"errors"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

func NewUserRepositoryInterface(DB *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		DB: DB,
	}
}

type userRepository struct {
	*gorm.DB
}

func (userRepository *userRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	user := entity.User{}
	result := userRepository.DB.WithContext(ctx).Where("user.email = ?", email).Find(&user)
	if result.RowsAffected == 0 {
		return entity.User{}, nil
	}

	return user, nil
}

func (userRepository *userRepository) FindByID(ctx context.Context, id int) (entity.User, error) {
	user := entity.User{}
	result := userRepository.DB.WithContext(ctx).Where("user.id = ?", id).Find(&user)
	if result.RowsAffected == 0 {
		return entity.User{}, nil
	}

	return user, nil
}

func (userRepository *userRepository) Insert(ctx context.Context, tx *gorm.DB, user entity.User) (int, error) {
	err := tx.WithContext(ctx).Create(&user).Error
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (userRepository *userRepository) FindByNoTelp(ctx context.Context, noTelp string) (entity.User, error) {
	user := entity.User{}
	result := userRepository.DB.WithContext(ctx).Where("user.notelp = ?", noTelp).Find(&user)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("User not found")
	}

	return user, nil
}

func (userRepository *userRepository) Update(ctx context.Context, tx *gorm.DB, user entity.User, id int) error {
	err := tx.WithContext(ctx).Where("user.id = ?", id).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}
