package delivery

import "myproject/innonformaledu/internal/features/class"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore class.ClassCore) ClassResponse {
	return ClassResponse{
		ID:                    dataCore.ID,
		ClassName:             dataCore.ClassName,
		ClassCategory:         dataCore.ClassCategory,
		ClassDescription:      dataCore.ClassDescription,
		ClassRule:             dataCore.ClassRule,
		ClassImage:            dataCore.ClassImage,
		AgencyCode:            dataCore.AgencyCode,
		AverageRating:         dataCore.AverageRating,
		SchoolAdministratorID: dataCore.SchoolAdministratorID,
		CreatedAt:             dataCore.CreatedAt,
		UpdatedAt:             dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []class.ClassCore) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
