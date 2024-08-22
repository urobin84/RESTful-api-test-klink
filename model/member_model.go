package model

type MemberRequest struct {
	Id   int16  `json:"id"`
	RecruiterId int16 `json:"recruiter_id"`
	FullName string `json:"full_name"`
	ActivationStatus string `json:"activation_status"`
}

type CreateMemberRequest struct {
	Id   int16  `json:"id"`
	RecruiterId int16 `json:"recruiter_id"`
	FullName string `json:"full_name"`
	ActivationStatus string `json:"activation_status"`
}

type SelectMemberRequest struct {
	Id   int16  `json:"id"`
	RecruiterId int16 `json:"recruiter_id"`
	FullName string `json:"full_name"`
	ActivationStatus string `json:"activation_status"`
}

type UpdateMemberRequest struct {
	Id   int16  `json:"id"`
	RecruiterId int16 `json:"recruiter_id"`
	FullName string `json:"full_name"`
	ActivationStatus string `json:"activation_status"`
}

type MemberResponse struct {
	Id   int16  `json:"id"`
	RecruiterId int16 `json:"recruiter_id"`
	FullName string `json:"full_name"`
	ActivationStatus string `json:"activation_status"`
}