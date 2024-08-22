package controller

import (
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(userService *service.ProductService) ProductController {
	return ProductController{
		ProductService: *userService,
	}
}

func (controller *ProductController) Route(app *fiber.App) {
	app.Get("/api/product", controller.SelectProduct)
	app.Post("/api/product", controller.CreateProduct)
	app.Put("/api/product", controller.UpdateProduct)
	app.Delete("/api/product", controller.DeleteProduct)
}

func (controller *ProductController) SelectProduct(c *fiber.Ctx) error {
	var request model.SelectProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.SelectProduct(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) CreateProduct(c *fiber.Ctx) error {
	var request model.CreateProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.CreateProduct(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) UpdateProduct(c *fiber.Ctx) error {
	var request model.UpdateProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.UpdateProduct(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *ProductController) DeleteProduct(c *fiber.Ctx) error {
	var request model.DeleteProductRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.ProductService.DeleteProduct(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}


