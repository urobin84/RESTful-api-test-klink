package entity

type Member struct {
	Entity
	RecruiterId				string `json:"recruiter_id" gorm:"not null"`
	MemberId				string `json:"member_id" gorm:"not null"`
	FullName			string `json:"full_name" gorm:"not null`
	Password			string `json:"password" gorm:"not null`
	Address			string `json:"address" gorm:"not null`
	Gender			string `json:"gender" gorm:"not null`
	IdentityNumber			string `json:"identity_number" gorm:"not null`
	Dob			string `json:"dob" gorm:"not null`
	Phone			string `json:"phone" gorm:"not null`
	Email			string `json:"email" gorm:"not null`
	Bank			string `json:"bank" gorm:"not null`
	Rekening			string `json:"rekening" gorm:"not null`
}