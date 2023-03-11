package repository

import "myproject/innonformaledu/internal/features/teacher"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

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
