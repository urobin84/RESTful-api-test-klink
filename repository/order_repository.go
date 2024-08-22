package repository

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
)

type OrderRepository interface {
	
	Insert(customer entity.Order) (response model.OrderResponse, err error)

	Update(customer entity.Order) (entity.Order, error)

	FindById(id string) (entity.Order, error)

	Find() []entity.Order

	DeleteAll()

	Delete(id string)

	Update(user entity.User) (entity.User, error)

	Delete(user entity.User) (entity.User, error)

	Select(user entity.User) (entity.User, error)

}
