package user

import "gorm.io/gorm"

func NewUserRepositoryInterface(DB *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		DB: DB,
	}
}

type userRepository struct {
	*gorm.DB
}
