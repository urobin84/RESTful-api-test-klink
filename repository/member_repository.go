package repository

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
)

type MemberRepository interface {
	Insert(customer entity.Member) (response model.MemberResponse, err error)

	Update(customer entity.Member) (entity.Member, error)

	FindByName(name string) (entity.Member, error)

	FindByEmail(email string) (entity.Member, error)

	FindById(id string) (entity.Member, error)

	Find() []entity.Member

	DeleteAll()

	Delete(id string)
}
