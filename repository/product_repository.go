package repository

import (
	"RESTful-api-test-klink/entity"
)

type ProductRepository interface {
	Insert(product entity.Product) (entity.Product, error)

	Update(product entity.Product) (entity.Product, error)

	FindByName(name string) (entity.Product, error)

	FindByNamePrice(name string, price int64) (entity.Product, error)

	FindById(id string) (entity.Product, error)

	FindAll() []entity.Product

	DeleteAll()

	Delete(id string)

	Upload(id string, images entity.ProductImage) (entity.ProductImage, error)

	
}
