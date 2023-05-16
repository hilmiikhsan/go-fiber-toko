package user

import (
	"context"

	"github.com/hilmiikhsan/go_rest_api/model"
)

type UserServiceInterface interface {
	GetProfile(ctx context.Context, userID int) (model.UserModel, error)
	UpdateProfile(ctx context.Context, user model.UpdateUserProfileModel, userID int) error
}
