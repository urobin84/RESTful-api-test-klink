package service

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/repository"
)

func NewOrderService(orderRepository *repository.OrderRepository) OrderService {
	return &OrderServiceImpl{
		repository: *orderRepository,
	}
}

type OrderServiceImpl struct {
	repository repository.OrderRepository
}

func (service OrderServiceImpl) Insert(order model.OrderRequest, user model.CreateUserResponse) (response model.OrderResponse, err error) {

	var orderInsert = entity.Order{
		OrderNumber:      order.OrderNumber,
	}
	return service.repository.Insert(orderInsert)
}

func (service OrderServiceImpl) Update(order model.OrderRequest) (entity.Order, error) {
	var orderUpdate = entity.Order{
		OrderNumber:      order.OrderNumber,
	}
	return service.repository.Update(orderUpdate)
}

func (service OrderServiceImpl) FindById(id string) (entity.Order, error) {
	return service.repository.FindById(id)
}

func (service OrderServiceImpl) SelectOrder() []entity.Order {
	return service.repository.Find()
}

func (service OrderServiceImpl) DeleteAll() {
	service.repository.DeleteAll()
}

func (service OrderServiceImpl) Delete(id string) {
	service.repository.Delete(id)
}
