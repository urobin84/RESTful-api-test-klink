package model

type ProductRequest struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty int `json:"qty"`
}

type GetProductResponse struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty int `json:"qty"`
}

type CreateProductRequest struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty int `json:"qty"`
}

type CreateProductResponse struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty int `json:"qty"`
	Images string `json:"image"`
}

type SelectProductRequest struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty int `json:"qty"`
	Images string `json:"image"`
}

type UpdateProductRequest struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty int `json:"qty"`
	Images string `json:"image"`
}

type ProductResponse struct {
	Id   int16  `json:"id"`
	Name string `json:"name"`
	Price string `json:"price"`
	Qty int `json:"qty"`
}

type ProductImage struct {
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
}

type UploadProductImageResponse struct {
	Id        string `json:"id"`
	ProductId string `json:"product_id"`
	Name      string `json:"name"`
	Path      string `json:"path"`
}