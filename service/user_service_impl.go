package service

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/exception"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/repository"
	"RESTful-api-test-klink/validation"
)

func NewUserService(userRepository *repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: *userRepository,
	}
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func (service userServiceImpl) Register(request model.CreateUserRequest) (user model.CreateUserResponse) {
	validation.ValidateUser(request)

	pass, err := validation.HashPassword(request.Password)
	exception.PanicIfNeeded(err)

	userRequest := entity.User{
		Username: request.Username,
		Password: pass,
		Email:    request.Email,
		Phone:    request.Phone,
	}

	err = service.repository.Insert(userRequest)
	exception.PanicIfNeeded(err)

	user = model.CreateUserResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return user
}

func (service userServiceImpl) FindByUsername(username string) (response model.CreateUserResponse) {
	user, err := service.repository.FindByUsername(username)
	exception.PanicIfNeeded(err)

	response = model.CreateUserResponse{
		// Id:       user.Id.String(),
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return response
}

func (service userServiceImpl) Login(request model.CreateUserRequest) (response model.GetLoginResponse) {
	validation.ValidateLogin(request)

	user, err := service.repository.Login(request.Email, request.Password)
	exception.PanicIfNeeded(err)

	token, err := validation.CreateToken(user)
	exception.PanicIfNeeded(err)

	response = model.GetLoginResponse{
		AccessToken: token,
		User: model.GetUserResponse{
			// Id:       user.Id.String(),
			Username: user.Username,
			// FullName: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			Role:     "admin",
		},
	}

	return response
}
