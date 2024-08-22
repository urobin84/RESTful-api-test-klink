package model

type InvoiceRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	TotalAmount string `json:"total_amount"`
	OrderId int16 `json:"order_id"`
	Status string `json:"status"`
	PaymentStatus string `json:"payment_status"`
}

type SelectInvoiceRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	TotalAmount string `json:"total_amount"`
	OrderId int16 `json:"order_id"`
	Status string `json:"status"`
	PaymentStatus string `json:"payment_status"`
}

type CreateInvoiceRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	TotalAmount string `json:"total_amount"`
	OrderId int16 `json:"order_id"`
	Status string `json:"status"`
	PaymentStatus string `json:"payment_status"`
}

type UpdateInvoiceRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	TotalAmount string `json:"total_amount"`
	OrderId int16 `json:"order_id"`
	Status string `json:"status"`
	PaymentStatus string `json:"payment_status"`
}


type InvoiceResponse struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Price string `json:"price"`
	Qty string `json:"qty"`
	TotalAmount string `json:"total_amount"`
	OrderId int16 `json:"order_id"`
	Status string `json:"status"`
	PaymentStatus string `json:"payment_status"`
}