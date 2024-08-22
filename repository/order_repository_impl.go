package repository

import (
	"gorm.io/gorm"
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
)

func NewOrderRepository(database *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{
		Database: database,
	}
}

type OrderRepositoryImpl struct {
	Database *gorm.DB
}

func (repository OrderRepositoryImpl) Insert(order entity.Order) (response model.OrderResponse, err error) {
	result := repository.Database.Create(&order).Scan(&order)
	if result.Error != nil {
		panic(result.Error)
	}

	response = model.OrderResponse{
		// ProductId: order.ProductId,
		// MemberId: order.MemberId,
		// RecruiterId: order.RecruiterId,
		// Status: order.Status,
		OrderNumber: order.OrderNumber,
		CreatedBy: order.CreatedBy,
		ProductSnapshot: order.ProductSnapshot,
		PaymentMethod: order.PaymentMethod,
		OrderStatus: order.OrderStatus,
	}

	return response, nil
}

func (repository OrderRepositoryImpl) Update(order entity.Order) (entity.Order, error) {
	var orderUpdate entity.Order
	repository.Database.Where("id = ?", order.Id).First(&orderUpdate)

	if len(orderUpdate.OrderNumber) == 0 {
		panic("Order not found")
	}

	repository.Database.Model(&orderUpdate).Updates(order)

	return orderUpdate, nil
}

func (repository OrderRepositoryImpl) FindById(id string) (entity.Order, error) {

	var order entity.Order
	repository.Database.Where("id = ?", id).First(&order)

	if len(order.OrderNumber) == 0 {
		panic("Order not found")
	}

	return order, nil
}

func (repository OrderRepositoryImpl) Find() []entity.Order {
	var orders []entity.Order
	repository.Database.Find(&orders)

	return orders
}

func (repository OrderRepositoryImpl) DeleteAll() {
	repository.Database.Delete(&entity.Order{})
}

func (repository OrderRepositoryImpl) Delete(id string) {
	repository.Database.Where("id = ?", id).Delete(&entity.Order{})
}
