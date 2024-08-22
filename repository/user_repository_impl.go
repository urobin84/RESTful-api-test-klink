package repository

import (
	"errors"
	"gorm.io/gorm"
	"RESTful-api-test-klink/config"
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/validation"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		Database: database,
	}
}

type userRepositoryImpl struct {
	Database *gorm.DB
}

func (repository userRepositoryImpl) Insert(user entity.User) error {
	database := config.GetDatabase()

	_, err := repository.FindByUsername(user.Username)
	if err == nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "USER_EXIST",
		})
	}

	result := database.Create(&user)
	if result.Error != nil {
		exception.PanicIfNeeded(exception.ValidationError{
			Message: "NAME_EXIST",
		})
	}

	return nil

}

func (repository userRepositoryImpl) FindByUsername(username string) (entity.User, error) {
	database := config.GetDatabase()

	var user entity.User
	database.Where("username = ?", username).First(&user)

	if len(user.Username) == 0 {
		return user, errors.New("USER_NOT_FOUND")
	}

	return user, nil
}

func (repository userRepositoryImpl) Login(email, password string) (entity.User, error) {
	database := config.GetDatabase()

	var user entity.User
	database.Where("email = ?", email).First(&user)

	match, _ := validation.ValidatePassword(password, user.Password)
	if !match {
		return user, errors.New("INVALID_PASSWORD")
	}

	return user, nil

}
