package model

type OrderRequest struct {
	Id   int16  `json:"id"`
	ProductId int16 `json:"product_id`
	MemberId int16 `json:"member_id"`
	RecruiterId string `json:"recruiter_id"`
	Status string `json:"status"`
	OrderNumber string `json:"order_number"`
	CreatedBy string `json:"created_by"`
	ProductSnapshot string `json:"product_snapshot`
	PaymentMethod string `json:"payment_method"`
	OrderStatus string `json:"order_status"`
}

type CreateOrderRequest struct {
	Id   int16  `json:"id"`
	ProductId int16 `json:"product_id`
	MemberId int16 `json:"member_id"`
	RecruiterId string `json:"recruiter_id"`
	Status string `json:"status"`
	OrderNumber string `json:"order_number"`
	CreatedBy string `json:"created_by"`
	ProductSnapshot string `json:"product_snapshot`
	PaymentMethod string `json:"payment_method"`
	OrderStatus string `json:"order_status"`
}

type SelectOrderRequest struct {
	Id   int16  `json:"id"`
	ProductId int16 `json:"product_id`
	MemberId int16 `json:"member_id"`
	RecruiterId string `json:"recruiter_id"`
	Status string `json:"status"`
	OrderNumber string `json:"order_number"`
	CreatedBy string `json:"created_by"`
	ProductSnapshot string `json:"product_snapshot`
	PaymentMethod string `json:"payment_method"`
	OrderStatus string `json:"order_status"`
}

type UpdateOrderRequest struct {
	Id   int16  `json:"id"`
	ProductId int16 `json:"product_id`
	MemberId int16 `json:"member_id"`
	RecruiterId string `json:"recruiter_id"`
	Status string `json:"status"`
	OrderNumber string `json:"order_number"`
	CreatedBy string `json:"created_by"`
	ProductSnapshot string `json:"product_snapshot`
	PaymentMethod string `json:"payment_method"`
	OrderStatus string `json:"order_status"`
}

type OrderResponse struct {
	Id   int16  `json:"id"`
	ProductId int16 `json:"product_id`
	MemberId int16 `json:"member_id"`
	RecruiterId string `json:"recruiter_id"`
	Status string `json:"status"`
	OrderNumber string `json:"order_number"`
	CreatedBy string `json:"created_by"`
	ProductSnapshot string `json:"product_snapshot`
	PaymentMethod string `json:"payment_method"`
	OrderStatus string `json:"order_status"`
}