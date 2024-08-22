package service

import (
	"RESTful-api-test-klink/model"
)

type MemberService interface {
	SelectMember(admin model.SelectMemberRequest) (response []model.MemberResponse)

	CreateMember(admin model.CreateMemberRequest) (response model.MemberResponse)

	UpdateMember(admin model.UpdateMemberRequest) (response model.MemberResponse)

	DeleteMember(id string)
}
