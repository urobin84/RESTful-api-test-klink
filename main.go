package main

import (
	"RESTful-api-test-klink/config"
	"RESTful-api-test-klink/controller"
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/repository"
	"RESTful-api-test-klink/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v2"
)


func main() {

	// Setup Configuration
	configuration := config.New()

	// use postgres db
	database := config.NewPostgresDatabase(configuration)

	// Setup Repository

	// admin repository
	// adminRepository := repository.NewAdminRepository(database)

	//invoice repository
	// MemberRepository := repository.NewMemberRepository(database)

	//product repository
	// productRepository := repository.NewProductRepository(database)

	//order repository
	// orderRepository := repository.NewOrderRepository(database)

	//invoice repository
	// InvoiceRepository := repository.NewInvoiceRepository(database)

	// Setup Service

	//member service
	// memberService := service.NewMemberService(&productRepository)
	
	//admin service
	// adminService := service.NewAdminService(&productRepository)
	
	//product service
	// productService := service.NewProductService(&productRepository)

	//order service
	// orderService := service.NewOrderService(&productRepository)

	//Invoice service
	// InvoiceService := service.NewInvoiceService(&productRepository)

	// Setup Controller

	//admin controller
	// adminController := controller.NewAdminController(&adminService)

	//member controller
	// memberController := controller.NewMemberController(&memberService)

	//product controller
	// productController := controller.NewProductController(&productService)

	//order controller
	// orderController := controller.NewOrderController(&orderService)

	//invoice controller
	// invoiceController := controller.NewInvoiceController(&invoiceService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(cors.New())
	app.Use(recover.New())

	file, err := os.OpenFile(configuration.Get("LOG_DIR"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	exception.PanicIfNeeded(err)
	defer file.Close()

	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path} ${body}\n",
		Output: file,
	}))

	// Setup Routing without Authorization
	userController.Route(app)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(configuration.Get("JWT_SECRET")),
	}))

	// Setup Routing with Authorizat
	// productController.Route(app)
	// adminController.Route(app)
	// memberController.Route(app)
	// orderController.Route(app)
	// invoiceController.Route(app)

	// Start App
	configuration.Get("APP_PORT")
	app.Listen(":" + configuration.Get("APP_PORT"))
	exception.PanicIfNeeded(err)

}