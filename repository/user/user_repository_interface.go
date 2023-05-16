package user

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/entity"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindByID(ctx context.Context, id int) (entity.User, error)
	Insert(ctx context.Context, tx *gorm.DB, user entity.User) (int, error)
	FindByNoTelp(ctx context.Context, noTelp string) (entity.User, error)
	Update(ctx context.Context, tx *gorm.DB, user entity.User, id int) error
}
