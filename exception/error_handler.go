package exception

import (
	"net/http"
	"RESTful-api-test-klink/model"

	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	_, ok := err.(ValidationError)
	if ok {
		return ctx.JSON(model.WebResponse{
			Code:   400,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	return ctx.JSON(model.WebResponse{
		Code:   500,
		Status: "INTERNAL_SERVER_ERROR",
		Data:   err.Error(),
	})
}

// Handle http 429 response
//
//	param c *fiber.Ctx
//	param data interface{}
//	param message string
//	return error
func TooManyRequestResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusTooManyRequests).JSON(fiber.Map{
		"error":   true,
		"message": "too many request",
	})
}
