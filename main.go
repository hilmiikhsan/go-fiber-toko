package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/controller/user"
	"github.com/hilmiikhsan/go_rest_api/exception"
	userRepo "github.com/hilmiikhsan/go_rest_api/repository/user"
	userService "github.com/hilmiikhsan/go_rest_api/service/user"
)

func main() {
	// setup configuration
	config := configuration.New()
	db := configuration.NewDatabase(config)
	redis := configuration.NewRedis(config)

	// repository
	userRepository := userRepo.NewUserRepositoryInterface(db)

	// rest client
	// httpBinRestClient := restclient.NewHttpBinRestClient()

	// service
	userService := userService.NewUserServiceInterface(&userRepository, redis)
	// httpBinService := httpbin.NewHttpBinServiceInterface(&httpBinRestClient)

	// controller
	_ = user.NewUserController(&userService, config)

	// setup fiber
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	// routing
	// userController.

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
