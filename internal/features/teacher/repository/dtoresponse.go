package repository

import "myproject/innonformaledu/internal/features/teacher"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

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
