package repository

import (
	"myproject/innonformaledu/features/teacher"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	ClassID    uint
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
	Portofolio string
	Expertise  string
	Material   string
}

func fromCore(dataCore teacher.TeacherCore) Teacher {
	modelData := Teacher{
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		Role:       dataCore.Role,
		Image:      dataCore.Image,
		AgencyCode: dataCore.AgencyCode,
		Gender:     dataCore.Gender,
		Address:    dataCore.Address,
		Portofolio: dataCore.Portofolio,
		Expertise:  dataCore.Expertise,
		Material:   dataCore.Material,
	}
	return modelData
}

func (dataModel *Teacher) toCore() teacher.TeacherCore {
	return teacher.TeacherCore{
		ClassID:    dataModel.ClassID,
		Name:       dataModel.Name,
		Email:      dataModel.Email,
		Password:   dataModel.Password,
		Role:       dataModel.Role,
		Image:      dataModel.Image,
		AgencyCode: dataModel.AgencyCode,
		Gender:     dataModel.Gender,
		Address:    dataModel.Address,
		Portofolio: dataModel.Portofolio,
		Expertise:  dataModel.Expertise,
		Material:   dataModel.Material,
		CreatedAt:  dataModel.CreatedAt,
		UpdatedAt:  dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []Teacher) []teacher.TeacherCore {
	var dataCore []teacher.TeacherCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
