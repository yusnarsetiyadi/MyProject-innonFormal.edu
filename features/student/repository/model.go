package repository

import (
	"myproject/innonformaledu/features/student"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
}

func fromCore(dataCore student.StudentCore) Student {
	modelData := Student{
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		Role:       dataCore.Role,
		Image:      dataCore.Image,
		AgencyCode: dataCore.AgencyCode,
		Gender:     dataCore.Gender,
		Address:    dataCore.Address,
	}
	return modelData
}

func (dataModel *Student) toCore() student.StudentCore {
	return student.StudentCore{
		ID:         dataModel.ID,
		Name:       dataModel.Name,
		Email:      dataModel.Email,
		Password:   dataModel.Password,
		Role:       dataModel.Role,
		Image:      dataModel.Image,
		AgencyCode: dataModel.AgencyCode,
		Gender:     dataModel.Gender,
		Address:    dataModel.Address,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []Student) []student.StudentCore {
	var dataCore []student.StudentCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
