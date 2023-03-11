package delivery

import "myproject/innonformaledu/internal/features/teacher"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore teacher.TeacherCore) TeacherResponse {
	return TeacherResponse{
		ClassID:    dataCore.ClassID,
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
		CreatedAt:  dataCore.CreatedAt,
		UpdatedAt:  dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []teacher.TeacherCore) []TeacherResponse {
	var dataResponse []TeacherResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
