package entity

type Invoice struct {
	Entity
	XId        string  `json:"xid" gorm:"not null"`
	Price				string `json:"price" gorm:"not null"`
	Qty				string `json:"qty" gorm:"not null"`
	TotalAmount				string `json:"total_amount" gorm:"not null"`
	Status				string `json:"status" gorm:"not null"`
	OrderId				string `json:"order_id" gorm:"not null"`
	InvoiceStatus				string `json:"invoice_status" gorm:"not null"`
}