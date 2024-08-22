package model

type UserRequest struct {
	Id   int16  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	
}

type CreateUserRequest struct {
	Id   int16  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type GetUserResponse struct {
	Id   int16  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role string `json:"role"`
}

type SelectUserRequest struct {
	Id   int16  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role string `json:"role"`
}

type UpdateUserRequest struct {
	Id   int16  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Role string `json:"role"`
}

type CreateUserResponse struct {
	Id   int16  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type UserResponse struct {
	Id   int16  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type GetLoginResponse struct {
	AccessToken string      `json:"access_token"`
	User        interface{} `json:"user"`
}