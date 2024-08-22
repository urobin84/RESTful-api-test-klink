package service

import (
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
	"RESTful-api-test-klink/repository"
)

func NewMemberService(memberRepository *repository.MemberRepository) MemberService {
	return &MemberServiceImpl{
		repository: *memberRepository,
	}
}

type MemberServiceImpl struct {
	repository repository.MemberRepository
}

func (service MemberServiceImpl) CreateMember(member model.MemberRequest, user model.CreateUserResponse) (response model.MemberResponse, err error) {

	var memberInsert = entity.Member{
		FullName:      member.FullName,
	}
	return service.repository.Insert(memberInsert)
}

func (service MemberServiceImpl) UpdateMember(member model.MemberRequest) (entity.Member, error) {
	var memberUpdate = entity.Member{
		FullName:      member.FullName,
	}
	return service.repository.UpdateMember(memberUpdate)
}

func (service MemberServiceImpl) FindByName(name string) (entity.Member, error) {
	return service.repository.FindByName(name)
}

func (service MemberServiceImpl) FindByEmail(email string) (entity.Member, error) {
	return service.repository.FindByEmail(email)
}

func (service MemberServiceImpl) FindById(id string) (entity.Member, error) {
	return service.repository.FindById(id)
}

func (service MemberServiceImpl) Find() []entity.Member {
	return service.repository.Find()
}

func (service MemberServiceImpl) DeleteMember(id string) {
	service.repository.Delete(id)
}
