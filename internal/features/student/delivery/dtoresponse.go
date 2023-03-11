package delivery

import "myproject/innonformaledu/internal/features/student"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore student.StudentCore) StudentResponse {
	return StudentResponse{
		ID:         dataCore.ID,
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		Role:       dataCore.Role,
		Image:      dataCore.Image,
		AgencyCode: dataCore.AgencyCode,
		Gender:     dataCore.Gender,
		Address:    dataCore.Address,
		CreatedAt:  dataCore.CreatedAt,
		UpdatedAt:  dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []student.StudentCore) []StudentResponse {
	var dataResponse []StudentResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
