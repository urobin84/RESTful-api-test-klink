package repository

import (
	"gorm.io/gorm"
	"RESTful-api-test-klink/config"
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/exception"
)

func NewAdminRepository(database *gorm.DB) AdminRepository {
	return &adminRepositoryImpl{
		Database: database,
	}
}

type adminRepositoryImpl struct {
	Database *gorm.DB
}

func (repository adminRepositoryImpl) Insert(admin entity.Admin) error {
	database := config.GetDatabase()

	// _, err := repository.FindByFullName(admin.FullName)
	// if err == nil {
	// 	exception.PanicIfNeeded(exception.ValidationError{
	// 		Message: "ADMIN_EXIST",
	// 	})
	// }

	result := database.Create(&admin)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "NAME_EXIST",
		})
	}

	return nil

}

func (repository adminRepositoryImpl) FindByName(fullName string) error {
	// database := config.GetDatabase()

	
	return nil

}



