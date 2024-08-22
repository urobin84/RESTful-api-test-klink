package repository

import (
	"gorm.io/gorm"
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
)

func NewInvoiceRepository(database *gorm.DB) InvoiceRepository {
	return &InvoiceRepositoryImpl{
		Database: database,
	}
}

type InvoiceRepositoryImpl struct {
	Database *gorm.DB
}

func (repository InvoiceRepositoryImpl) Insert(invoice entity.Invoice) (response model.InvoiceResponse, err error) {
	result := repository.Database.Create(&invoice).Scan(&invoice)
	if result.Error != nil {
		panic(result.Error)
	}

	response = model.InvoiceResponse{
		
	}

	return response, nil
}

func (repository InvoiceRepositoryImpl) Update(invoice entity.Invoice) (entity.Invoice, error) {
	var invoiceUpdate entity.Invoice
	repository.Database.Where("id = ?", invoice.Id).First(&invoiceUpdate)

	// if len(invoiceUpdate.Id) == 0 {
	// 	panic("Invoice not found")
	// }

	repository.Database.Model(&invoiceUpdate).Updates(invoice)

	return invoiceUpdate, nil
}

func (repository InvoiceRepositoryImpl) FindById(id string) (entity.Invoice, error) {

	var invoice entity.Invoice
	repository.Database.Where("id = ?", id).First(&invoice)

	if len(invoice.XId) == 0 {
		panic("Invoice not found")
	}

	return invoice, nil
}

func (repository InvoiceRepositoryImpl) Find() []entity.Invoice {
	var invoices []entity.Invoice
	repository.Database.Find(&invoices)

	return invoices
}

func (repository InvoiceRepositoryImpl) DeleteAll() {
	repository.Database.Delete(&entity.Invoice{})
}

func (repository InvoiceRepositoryImpl) Delete(id string) {
	repository.Database.Where("id = ?", id).Delete(&entity.Invoice{})
}
