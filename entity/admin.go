package entity

type Admin struct {
	Entity
	XId				string `json:"xid" gorm:"not null"`
	Email			string `json:"email" gorm:"not null`
	Phone			string `json:"phone" gorm:"not null`
	FullName			string `json:"full_name" gorm:"not null`
	Role			string `json:"role" gorm:"not null`
}