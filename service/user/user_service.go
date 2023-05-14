package user

import (
	"github.com/go-redis/redis/v8"
	"github.com/hilmiikhsan/go_rest_api/repository/user"
)

func NewUserServiceInterface(userRepository *user.UserRepositoryInterface, cache *redis.Client) UserServiceInterface {
	return &userService{
		UserRepositoryInterface: *userRepository,
		Cache:                   cache,
	}
}

type userService struct {
	user.UserRepositoryInterface
	Cache *redis.Client
}
