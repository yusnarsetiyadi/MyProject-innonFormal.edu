package delivery

import "myproject/innonformaledu/internal/features/schooladministrator"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore schooladministrator.SchoolAdministratorCore) SchoolAdministratorResponse {
	return SchoolAdministratorResponse{
		ID:              dataCore.ID,
		Name:            dataCore.Name,
		Email:           dataCore.Email,
		Password:        dataCore.Password,
		Role:            dataCore.Role,
		Image:           dataCore.Image,
		AgencyCode:      dataCore.AgencyCode,
		NumberSchool:    dataCore.NumberSchool,
		AgreementLetter: dataCore.AgreementLetter,
		SuperAdminID:    dataCore.SuperAdminID,
		CreatedAt:       dataCore.CreatedAt,
		UpdatedAt:       dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []schooladministrator.SchoolAdministratorCore) []SchoolAdministratorResponse {
	var dataResponse []SchoolAdministratorResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
