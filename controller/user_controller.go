package controller

import (
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService *service.UserService) UserController {
	return UserController{
		UserService: *userService,
	}
}

func (controller *UserController) Route(app *fiber.App) {
	app.Post("/api/signup", controller.Register)
	app.Post("/api/login", controller.Login)
}

func (controller *UserController) Register(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.UserService.Register(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var request model.CreateUserRequest
	err := c.BodyParser(&request)

	exception.PanicIfNeeded(err)

	response := controller.UserService.Login(request)

	//Create Cookie
	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = response.AccessToken
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HTTPOnly = false
	cookie.Secure = true
	cookie.SameSite = "Strict"
	
	c.Cookie(cookie)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}
