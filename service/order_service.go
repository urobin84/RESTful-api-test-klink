package service

import (
	"RESTful-api-test-klink/model"
)

type OrderService interface {
	SelectOrder(admin model.SelectOrderRequest) (response []model.OrderResponse)

	CreateOrder(admin model.CreateOrderRequest) (response model.OrderResponse)

	UpdateOrder(admin model.UpdateOrderRequest) (response model.OrderResponse)

	DeleteOrder(id string)
}
