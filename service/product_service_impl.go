package service

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/repository"
	"RESTful-api-test-klink/validation"
)

func NewProductService(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		repository: *productRepository,
	}
}

type productServiceImpl struct {
	repository repository.ProductRepository
}

func (service *productServiceImpl) CreateProduct(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.ValidateProduct(request)

	product := entity.Product{
		Name:     request.Name,
		Price:    request.Price,
		// Quantity: request.Quantity,
	}

	service.repository.Insert(product)

	response = model.CreateProductResponse{
		// Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		// Quantity: product.Quantity,
	}
	return response
}

func (service *productServiceImpl) UpdateProduct(request model.CreateProductRequest) (response model.CreateProductResponse) {
	validation.ValidateProduct(request)

	product := entity.Product{
		Name:     request.Name,
		Price:    request.Price,
		// Quantity: request.Quantity,
	}

	service.repository.UpdateProduct(product)

	response = model.CreateProductResponse{
		// Id:       product.Id,
		Name:     product.Name,
		Price:    product.Price,
		// Quantity: product.Quantity,
	}
	return response
}

func (service *productServiceImpl) Delete(id string) {
	service.repository.Delete(id)
}

func (service *productServiceImpl) Find() (responses []model.GetProductResponse) {
	products := service.repository.FindAll()

	for _, product := range products {
		// var images []model.ProductImage
		// for _, image := range product.Images {
		// 	images = append(images, model.ProductImage{
		// 		Name: image.Name,
		// 		Path: image.Path,
		// 	})
		// }

		responses = append(responses, model.GetProductResponse{
			// Id:       product.Id,
			Name:     product.Name,
			Price:    product.Price,
			// Quantity: product.Quantity,
		})
	}
	return responses
}

func (service *productServiceImpl) FindById(id string) (response model.CreateProductResponse) {
	product, err := service.repository.FindById(id)
	if err != nil {
		panic(err)
	}
	// var images []model.ProductImage

	// for _, image := range product.Images {
	// 	images = append(images, model.ProductImage{
	// 		Name: image.Name,
	// 		Path: image.Path,
	// 	})
	// }

	return model.CreateProductResponse{
		// Id:       product.Id.String(),
		Name:     product.Name,
		Price:    product.Price,
		Qty: 			product.Qty,
		// Images:   images,
	}
}

func (service *productServiceImpl) Upload(id string, image model.ProductImage) (response model.UploadProductImageResponse) {
	var productImage entity.ProductImage
	productImage = entity.ProductImage{
		ProductID: id,
		Name:      image.Name,
		Path:      image.Path,
	}
	upload, err := service.repository.Upload(id, productImage)
	if err != nil {
		return model.UploadProductImageResponse{}
	}

	return model.UploadProductImageResponse{
		Id:        upload.ProductID,
		ProductId: id,
		Name:      upload.Name,
		Path:      upload.Path,
	}

}
