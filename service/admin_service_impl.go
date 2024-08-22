package service

import (
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/repository"
)

func NewAdminService(adminRepository *repository.AdminRepository) AdminService {
	return &adminServiceImpl{
		repository: *adminRepository,
	}
}

type adminServiceImpl struct {
	repository repository.AdminRepository
}

func (service adminServiceImpl) FindByName(name string) (response model.CreateAdminResponse) {
	// admin := service.repository.FindByName(name)

	response = model.CreateAdminResponse{
		// Id:       admin.Id.String(),
		// FullName: admin.FullName,
		// Email:    admin.Email,
		// Phone:    admin.Phone,
	}

	return response
}


