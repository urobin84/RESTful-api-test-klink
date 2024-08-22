package controller

import (
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type InvoiceController struct {
	InvoiceService service.InvoiceService
}

func NewInvoiceController(userService *service.InvoiceService) InvoiceController {
	return InvoiceController{
		InvoiceService: *userService,
	}
}

func (controller *InvoiceController) Route(app *fiber.App) {
	app.Get("/api/invoice", controller.SelectInvoice)
	app.Post("/api/invoice", controller.CreateInvoice)
	app.Put("/api/invoice", controller.UpdateInvoice)
	app.Delete("/api/invoice", controller.DeleteInvoice)
}

func (controller *InvoiceController) SelectInvoice(c *fiber.Ctx) error {
	var request model.CreateInvoiceRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.InvoiceService.SelectInvoice(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *InvoiceController) CreateInvoice(c *fiber.Ctx) error {
	var request model.CreateInvoiceRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.InvoiceService.CreateInvoice(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *InvoiceController) UpdateInvoice(c *fiber.Ctx) error {
	var request model.CreateInvoiceRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.InvoiceService.UpdateInvoice(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *InvoiceController) DeleteInvoice(c *fiber.Ctx) error {
	var request model.CreateInvoiceRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.InvoiceService.DeleteInvoice(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
