package entity

type Product struct {
	Entity
	Code				string `json:"code" gorm:"not null"`
	Name				string `json:"name" gorm:"not null"`
	Price				string `json:"price" gorm:"not null"`
	Qty			int `json:"qty" gorm:"not null`
	Description			string `json:"description" gorm:"not null`
	Image			string `json:"image" gorm:"not null`
}

type ProductImage struct {
	Entity
	ProductID string `json:"product_id" gorm:"not null"`
	Name      string `json:"name" gorm:"max:255;not null;unique"`
	Path      string `json:"path" gorm:"max:255;null"`
}