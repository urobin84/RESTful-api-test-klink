package model

type AdminRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
}

type CreateAdminRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
}

type UpdateAdminRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
}

type SelectAdminRequest struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
}

type CreateAdminResponse struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
}

type AdminResponse struct {
	Id   int16  `json:"id"`
	XId string `json:"xid"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	FullName string `json:"full_name"`
	Role string `json:"role"`
}