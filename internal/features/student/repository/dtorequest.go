package repository

import "myproject/innonformaledu/internal/features/student"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

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
