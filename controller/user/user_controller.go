package user

import (
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/service/user"
)

type UserController struct {
	user.UserServiceInterface
	configuration.Config
}

func NewUserController(userService *user.UserServiceInterface, config configuration.Config) *UserController {
	return &UserController{
		UserServiceInterface: *userService,
		Config:               config,
	}
}
