package repository

import (
	"errors"
	"gorm.io/gorm"
	"RESTful-api-test-klink/config"
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/exception"
)

func NewProductRepository(database *gorm.DB) ProductRepository {
	return &productRepositoryImpl{
		Database: database,
	}
}

type productRepositoryImpl struct {
	Database *gorm.DB
}

func (repository productRepositoryImpl) Insert(product entity.Product) (entity.Product, error) {

	database := config.GetDatabase()

	_, err := repository.FindByName(product.Name)
	if err != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "PRODUCT_ALREADY_EXISTS",
		})
	}

	result := database.Create(&product)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: result.Error.Error(),
		})
	}

	return product, nil

}

func (repository productRepositoryImpl) Update(product entity.Product) (entity.Product, error) {

	database := config.GetDatabase()
	var productUpdate entity.Product
	database.Where("id = ?", product.Id).First(&productUpdate)

	if len(productUpdate.Name) == 0 {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "PRODUCT_NOT_FOUND",
		})

	}

	database.Model(&productUpdate).Updates(product)

	return productUpdate, nil

}

func (repository productRepositoryImpl) FindByName(name string) (entity.Product, error) {

	database := config.GetDatabase()

	var product entity.Product
	database.Where("name = ?", name).First(&product)

	return product, nil
}

func (repository productRepositoryImpl) FindByNamePrice(name string, price int64) (entity.Product, error) {
	database := config.GetDatabase()

	var product entity.Product
	database.Where("name = ? and price = ?", name, price).Preload("Product").First(&product)

	if len(product.Name) == 0 {
		return product, errors.New("PRODUCT_NOT_FOUND")
	}

	return product, nil
}

func (repository productRepositoryImpl) FindById(id string) (entity.Product, error) {

	database := config.GetDatabase()

	var product entity.Product
	database.Where("id = ?", id).Preload("Product").First(&product)

	if len(product.Name) == 0 {
		return product, errors.New("PRODUCT_NOT_FOUND")
	}

	return product, nil
}

func (repository productRepositoryImpl) FindAll() (products []entity.Product) {

	database := config.GetDatabase()

	database.Find(&products)

	return products
}

func (repository productRepositoryImpl) DeleteAll() {

	database := config.GetDatabase()

	database.Delete(&entity.Product{})
}

func (repository productRepositoryImpl) Delete(id string) {

	database := config.GetDatabase()

	database.Where("id = ?", id).Delete(&entity.Product{})
}

func (repository productRepositoryImpl) Upload(id string, image entity.ProductImage) (entity.ProductImage, error) {

	database := config.GetDatabase()

	var productUpdate entity.Product
	database.Where("id = ?", id).First(&productUpdate)

	if len(productUpdate.Name) == 0 {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "PRODUCT_NOT_FOUND",
		})
	}

	result := database.Create(&image)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: result.Error.Error(),
		})
	}

	return image, nil
}
