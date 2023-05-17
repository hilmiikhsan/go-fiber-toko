package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/controller/auth"
	"github.com/hilmiikhsan/go_rest_api/controller/province"
	"github.com/hilmiikhsan/go_rest_api/controller/toko"
	"github.com/hilmiikhsan/go_rest_api/controller/user"
	"github.com/hilmiikhsan/go_rest_api/exception"
	alamatRepo "github.com/hilmiikhsan/go_rest_api/repository/alamat"
	tokoRepo "github.com/hilmiikhsan/go_rest_api/repository/toko"
	userRepo "github.com/hilmiikhsan/go_rest_api/repository/user"
	alamatService "github.com/hilmiikhsan/go_rest_api/service/alamat"
	authService "github.com/hilmiikhsan/go_rest_api/service/auth"
	provinceService "github.com/hilmiikhsan/go_rest_api/service/province"
	tokoService "github.com/hilmiikhsan/go_rest_api/service/toko"
	userService "github.com/hilmiikhsan/go_rest_api/service/user"
)

func main() {
	// setup configuration
	config := configuration.New()
	db := configuration.NewDatabase(config)
	// redis := configuration.NewRedis(config)

	// repository
	userRepository := userRepo.NewUserRepositoryInterface(db)
	tokoRepository := tokoRepo.NewTokoRepositoryInterface(db)
	alamatRepository := alamatRepo.NewAlamatRepositoryInterface(db)

	// rest client
	// httpBinRestClient := restclient.NewHttpBinRestClient()

	// service
	authService := authService.NewAuthServiceInterface(&userRepository, &tokoRepository, db)
	userService := userService.NewUserServiceInterface(&userRepository, db)
	alamatService := alamatService.NewAlamatServiceInterface(&alamatRepository, db)
	tokoService := tokoService.NewTokoServiceInterface(&tokoRepository, db)
	provinceService := provinceService.NewProvinceServiceInterface()
	// httpBinService := httpbin.NewHttpBinServiceInterface(&httpBinRestClient)

	// controller
	authController := auth.NewAuthController(&authService, config)
	userController := user.NewUserController(&userService, config, &alamatService)
	tokoController := toko.NewTokoController(&tokoService, config)
	provinceController := province.NewProvinceController(&provinceService, config)

	// setup fiber
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	// routing
	authController.Route(app)
	userController.Route(app)
	tokoController.Route(app)
	provinceController.Route(app)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
