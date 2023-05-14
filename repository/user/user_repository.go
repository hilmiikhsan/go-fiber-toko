package user

import (
	"context"

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

func (userRepository *userRepository) RegisterUser(ctx context.Context, user entity.User) error {
	err := userRepository.DB.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}
