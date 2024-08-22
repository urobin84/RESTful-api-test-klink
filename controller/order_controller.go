package controller

import (
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type OrderController struct {
	OrderService service.OrderService
}

func NewOrderController(userService *service.OrderService) OrderController {
	return OrderController{
		OrderService: *userService,
	}
}

func (controller *OrderController) Route(app *fiber.App) {
	app.Get("/api/order", controller.SelectOrder)
	app.Post("/api/order", controller.CreateOrder)
	app.Put("/api/order", controller.UpdateOrder)
	app.Delete("/api/order", controller.DeleteOrder)
}

func (controller *OrderController) SelectOrder(c *fiber.Ctx) error {
	var request model.SelectOrderRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.OrderService.SelectOrder(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *OrderController) CreateOrder(c *fiber.Ctx) error {
	var request model.CreateOrderRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.OrderService.CreateOrder(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *OrderController) UpdateOrder(c *fiber.Ctx) error {
	var request model.UpdateOrderRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.OrderService.UpdateOrder(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *OrderController) DeleteOrder(c *fiber.Ctx) error {
	var request model.DeleteOrderRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.OrderService.DeleteOrder(request)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
