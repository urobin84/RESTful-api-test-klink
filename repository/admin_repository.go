package repository

import "RESTful-api-test-klink/entity"

type AdminRepository interface {
	Insert(admin entity.Admin) error
	FindByName(admin string) error
}
