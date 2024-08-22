package controller

import (
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MemberController struct {
	MemberService service.MemberService
}

func NewMemberController(userService *service.MemberService) MemberController {
	return MemberController{
		MemberService: *userService,
	}
}

func (controller *MemberController) Route(app *fiber.App) {
	app.GET("/api/member", controller.SelectMember)
	app.Post("/api/member", controller.CreateMember)
	app.Put("/api/member", controller.UpdateMember)
	app.Delete("/api/member", controller.DeleteMember)
}

func (controller *MemberController) SelectMember(c *fiber.Ctx) error {
	var request model.SelectMemberRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.MemberService.SelectMember(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberController) CreateMember(c *fiber.Ctx) error {
	var request model.CreateMemberRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.MemberService.CreateMember(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberController) UpdateMember(c *fiber.Ctx) error {
	var request model.UpdateMemberRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.MemberService.UpdateMember(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *MemberController) DeleteMember(c *fiber.Ctx) error {
	var request model.UpdateMemberRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.MemberService.DeleteMember(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}


