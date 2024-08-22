package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
)

func ValidateAdmin(request model.CreateAdminRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.XId, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.Phone, validation.Required),
		validation.Field(&request.FullName, validation.Required),
		validation.Field(&request.Role, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
