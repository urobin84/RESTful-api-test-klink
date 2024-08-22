package repository

import "RESTful-api-test-klink/entity"

type UserRepository interface {
	
	Insert(user entity.User) error

	FindByUsername(username string) (entity.User, error)

	Login(email, password string) (entity.User, error)

	Update(user entity.User) (entity.User, error)

	Delete(user entity.User) (entity.User, error)

	Select(user entity.User) (entity.User, error)

}
