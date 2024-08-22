package repository

import (
	"gorm.io/gorm"
	"RESTful-api-test-klink/entity"
	"RESTful-api-test-klink/model"
)

func NewMemberRepository(database *gorm.DB) MemberRepository {
	return &MemberRepositoryImpl{
		Database: database,
	}
}

type MemberRepositoryImpl struct {
	Database *gorm.DB
}

func (repository MemberRepositoryImpl) Insert(member entity.Member) (response model.MemberResponse, err error) {
	result := repository.Database.Create(&member).Scan(&member)
	if result.Error != nil {
		panic(result.Error)
	}

	response = model.MemberResponse{
		// Id:	 member.id,
		FullName:  member.FullName,
	}

	return response, nil
}

func (repository MemberRepositoryImpl) Update(member entity.Member) (entity.Member, error) {
	var memberUpdate entity.Member
	repository.Database.Where("id = ?", member.Id).First(&memberUpdate)

	if len(memberUpdate.FullName) == 0 {
		panic("Member not found")
	}

	repository.Database.Model(&memberUpdate).Updates(member)

	return memberUpdate, nil
}

func (repository MemberRepositoryImpl) FindByName(name string) (entity.Member, error) {

	var member entity.Member
	repository.Database.Where("name = ?", name).First(&member)

	if len(member.FullName) == 0 {
		panic("Member not found")
	}

	return member, nil
}

func (repository MemberRepositoryImpl) FindByEmail(email string) (entity.Member, error) {

	var member entity.Member
	repository.Database.Where("email = ?", email).First(&member)

	if len(member.FullName) == 0 {
		panic("Member not found")
	}

	return member, nil
}

func (repository MemberRepositoryImpl) FindById(id string) (entity.Member, error) {

	var member entity.Member
	repository.Database.Where("id = ?", id).First(&member)

	if len(member.FullName) == 0 {
		panic("Member not found")
	}

	return member, nil
}

func (repository MemberRepositoryImpl) Find() []entity.Member {
	var members []entity.Member
	repository.Database.Find(&members)

	return members
}

func (repository MemberRepositoryImpl) DeleteAll() {
	repository.Database.Delete(&entity.Member{})
}

func (repository MemberRepositoryImpl) Delete(id string) {
	repository.Database.Where("id = ?", id).Delete(&entity.Member{})
}
