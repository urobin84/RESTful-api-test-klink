package service

import (
	"RESTful-api-test-klink/model"
)

type InvoiceService interface {
	SelectInvoice(admin model.SelectInvoiceRequest) (response []model.InvoiceResponse)

	CreateInvoice(admin model.CreateInvoiceRequest) (response model.InvoiceResponse)

	UpdateInvoice(admin model.UpdateInvoiceRequest) (response model.InvoiceResponse)

	DeleteInvoice(id string)
}
