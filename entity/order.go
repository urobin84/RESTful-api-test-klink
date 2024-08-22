package entity

type Order struct {
	Entity
	OrderNumber				string `json:"order_number" gorm:"not null"`
	RecriuterId				string `json:"recriuter_id" gorm:"not null"`
	ProductSnapshot				string `json:"product_snapshot" gorm:"not null"`
	MemberSnapshot				string `json:"member_snapshot" gorm:"not null"`
	PaymentMethod				string `json:"payment_method" gorm:"not null"`
	OrderStatus				string `json:"OrderStatus" gorm:"not null"`
	CreatedBy				string `json:"created_by" gorm:"not null"`
}