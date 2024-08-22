package service

import (
	"RESTful-api-test-klink/model"
)

type AdminService interface {

	FindByName(name string) (admin model.CreateAdminResponse)

	SelectAdmin(admin model.SelectAdminRequest) (response []model.AdminResponse)

	CreateAdmin(admin model.CreateAdminRequest) (response model.AdminResponse)

	UpdateAdmin(admin model.UpdateAdminRequest) (response model.AdminResponse)

	DeleteAdmin(id string)

}
