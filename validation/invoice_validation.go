package validation

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	// "github.com/go-ozzo/ozzo-validation/v4/is"
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
)

func ValidateInvoice(request model.CreateInvoiceRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.XId, validation.Required),
		validation.Field(&request.Price, validation.Required),
		validation.Field(&request.Qty, validation.Required),
		validation.Field(&request.TotalAmount, validation.Required),
		validation.Field(&request.OrderId, validation.Required),
		validation.Field(&request.Status, validation.Required),
		validation.Field(&request.PaymentStatus, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
