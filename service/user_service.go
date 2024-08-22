package service

import (
	"RESTful-api-test-klink/model"
)

type UserService interface {
	Register(request model.CreateUserRequest) (response model.CreateUserResponse)

	FindByUsername(username string) (user model.CreateUserResponse)

	Login(request model.CreateUserRequest) (response model.GetLoginResponse)

	SelectUser(user model.SelectUserRequest) (response []model.UserResponse)

	CreateUser(user model.CreateUserRequest) (response model.UserResponse)

	UpdateUser(user model.UpdateUserRequest) (response model.UserResponse)
	
	DeleteUser(id string)
}
