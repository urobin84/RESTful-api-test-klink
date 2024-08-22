package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/go-ozzo/ozzo-validation/v4/is"
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
)

func ValidateOrder(request model.CreateOrderRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.ProductId, validation.Required),
		validation.Field(&request.MemberId, validation.Required),
		validation.Field(&request.RecruiterId, validation.Required),
		validation.Field(&request.Status, validation.Required),
		validation.Field(&request.OrderNumber, validation.Required),
		validation.Field(&request.CreatedBy, validation.Required),
		validation.Field(&request.ProductSnapshot, validation.Required),
		validation.Field(&request.PaymentMethod, validation.Required),
		validation.Field(&request.OrderStatus, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
