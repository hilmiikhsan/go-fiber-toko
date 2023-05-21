package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hilmiikhsan/go_rest_api/configuration"
	"github.com/hilmiikhsan/go_rest_api/controller/auth"
	"github.com/hilmiikhsan/go_rest_api/controller/category"
	"github.com/hilmiikhsan/go_rest_api/controller/product"
	"github.com/hilmiikhsan/go_rest_api/controller/province"
	"github.com/hilmiikhsan/go_rest_api/controller/toko"
	"github.com/hilmiikhsan/go_rest_api/controller/user"
	"github.com/hilmiikhsan/go_rest_api/exception"
	alamatRepo "github.com/hilmiikhsan/go_rest_api/repository/alamat"
	categoryRepo "github.com/hilmiikhsan/go_rest_api/repository/category"
	fotoProdukRepo "github.com/hilmiikhsan/go_rest_api/repository/foto_produk"
	productRepo "github.com/hilmiikhsan/go_rest_api/repository/product"
	tokoRepo "github.com/hilmiikhsan/go_rest_api/repository/toko"
	userRepo "github.com/hilmiikhsan/go_rest_api/repository/user"
	alamatService "github.com/hilmiikhsan/go_rest_api/service/alamat"
	authService "github.com/hilmiikhsan/go_rest_api/service/auth"
	categoryService "github.com/hilmiikhsan/go_rest_api/service/category"
	productService "github.com/hilmiikhsan/go_rest_api/service/product"
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
	categoryRepository := categoryRepo.NewCategoryRepositoryInterface(db)
	productRepository := productRepo.NewProductRepositoryInterface(db)
	fotoProdukRepository := fotoProdukRepo.NewFotoProdukRepositoryInterface(db)

	// rest client
	// httpBinRestClient := restclient.NewHttpBinRestClient()

	// service
	authService := authService.NewAuthServiceInterface(&userRepository, &tokoRepository, db)
	provinceService := provinceService.NewProvinceServiceInterface()
	userService := userService.NewUserServiceInterface(&userRepository, db, &provinceService)
	alamatService := alamatService.NewAlamatServiceInterface(&alamatRepository, db)
	tokoService := tokoService.NewTokoServiceInterface(&tokoRepository, db)
	categoryService := categoryService.NewCategoryServiceInterface(&categoryRepository, db, &userRepository)
	productService := productService.NewProductServiceInterface(&productRepository, db, &tokoRepository, &fotoProdukRepository)
	// httpBinService := httpbin.NewHttpBinServiceInterface(&httpBinRestClient)

	// controller
	authController := auth.NewAuthController(&authService, config)
	userController := user.NewUserController(&userService, config, &alamatService)
	tokoController := toko.NewTokoController(&tokoService, config)
	provinceController := province.NewProvinceController(&provinceService, config)
	categoryController := category.NewCategoryController(&categoryService, config)
	productController := product.NewProductController(&productService, config)

	// setup fiber
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New())

	// routing
	authController.Route(app)
	userController.Route(app)
	tokoController.Route(app)
	provinceController.Route(app)
	categoryController.Route(app)
	productController.Route(app)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
