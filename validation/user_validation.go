package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
)

func ValidateUser(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Username, validation.Required),
		validation.Field(&request.Password, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateLogin(request model.CreateUserRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.Password, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
