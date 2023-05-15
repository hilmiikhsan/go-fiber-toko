package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/controller/auth"
	"github.com/hilmiikhsan/go_rest_api/controller/user"
	"github.com/hilmiikhsan/go_rest_api/exception"
	userRepo "github.com/hilmiikhsan/go_rest_api/repository/user"
	authService "github.com/hilmiikhsan/go_rest_api/service/auth"
	userService "github.com/hilmiikhsan/go_rest_api/service/user"
)

func main() {
	// setup configuration
	config := configuration.New()
	db := configuration.NewDatabase(config)
	// redis := configuration.NewRedis(config)

	// repository
	userRepository := userRepo.NewUserRepositoryInterface(db)

	// rest client
	// httpBinRestClient := restclient.NewHttpBinRestClient()

	// service
	authService := authService.NewAuthServiceInterface(&userRepository)
	userService := userService.NewUserServiceInterface(&userRepository)
	// httpBinService := httpbin.NewHttpBinServiceInterface(&httpBinRestClient)

	// controller
	authController := auth.NewAuthController(&authService, config)
	userController := user.NewUserController(&userService, config)

	// setup fiber
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	// routing
	authController.Route(app)
	userController.Route(app)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
