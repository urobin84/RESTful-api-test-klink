package controller

import (
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AdminController struct {
	AdminService service.AdminService
}

func NewAdminController(userService *service.AdminService) AdminController {
	return AdminController{
		AdminService: *userService,
	}
}

func (controller *AdminController) Route(app *fiber.App) {
	app.Get("/api/admin", controller.SelectAdmin)
	app.Post("/api/admin", controller.CreateAdmin)
	app.Put("/api/admin", controller.UpdateAdmin)
	app.Delete("/api/admin", controller.DeleteAdmin)
}

func (controller *AdminController) SelectAdmin(c *fiber.Ctx) error {
	var request model.CreateAdminRequest
	// err := c.BodyParser(&request)

	response := controller.AdminService.SelectAdmin(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AdminController) DeleteAdmin(c *fiber.Ctx) error {
	var request model.CreateAdminRequest
	// err := c.BodyParser(&request)

	response := controller.AdminService.DeleteAdmin(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AdminController) CreateAdmin(c *fiber.Ctx) error {
	var request model.CreateAdminRequest
	// err := c.BodyParser(&request)

	response := controller.AdminService.CreateAdmin(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *AdminController) UpdateAdmin(c *fiber.Ctx) error {
	var request model.CreateAdminRequest
	// err := c.BodyParser(&request)

	response := controller.AdminService.UpdateAdmin(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
