package repository

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
)

type InvoiceRepository interface {
	Insert(customer entity.Invoice) (response model.InvoiceResponse, err error)

	Update(customer entity.Invoice) (entity.Invoice, error)

	FindById(id string) (entity.Invoice, error)

	Find() []entity.Invoice

	DeleteAll()

	Delete(id string)
}
