package service

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/repository"
)

func NewInvoiceService(invoiceRepository *repository.InvoiceRepository) InvoiceService {
	return &InvoiceServiceImpl{
		repository: *invoiceRepository,
	}
}

type InvoiceServiceImpl struct {
	repository repository.InvoiceRepository
}

func (service InvoiceServiceImpl) Insert(invoice model.InvoiceRequest, user model.CreateUserResponse) (response model.InvoiceResponse, err error) {

	var invoiceInsert = entity.Invoice{
		XId:      invoice.XId,
	}
	return service.repository.Insert(invoiceInsert)
}

func (service InvoiceServiceImpl) Update(invoice model.InvoiceRequest) (entity.Invoice, error) {
	var invoiceUpdate = entity.Invoice{
		XId:      invoice.XId,
	}
	return service.repository.Update(invoiceUpdate)
}

func (service InvoiceServiceImpl) FindById(id string) (entity.Invoice, error) {
	return service.repository.FindById(id)
}

func (service InvoiceServiceImpl) Find() []entity.Invoice {
	return service.repository.Find()
}

func (service InvoiceServiceImpl) DeleteAll() {
	service.repository.DeleteAll()
}

func (service InvoiceServiceImpl) Delete(id string) {
	service.repository.Delete(id)
}
