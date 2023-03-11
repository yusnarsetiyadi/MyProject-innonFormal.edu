package repository

import "myproject/innonformaledu/internal/features/student"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

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
