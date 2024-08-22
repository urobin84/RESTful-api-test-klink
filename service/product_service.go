package service

import "RESTful-api-test-klink/model"

type ProductService interface {

	SelectProduct(product model.SelectProductRequest) (response []model.ProductResponse)

	CreateProduct(product model.CreateProductRequest) (response model.ProductResponse)

	UpdateProduct(product model.UpdateProductRequest) (response model.ProductResponse)

	DeleteProduct(id string)
	
}
