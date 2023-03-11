package repository

import "myproject/innonformaledu/internal/features/schooladministrator"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

func fromCore(dataCore schooladministrator.SchoolAdministratorCore) SchoolAdministrator {
	modelData := SchoolAdministrator{
		Name:            dataCore.Name,
		Email:           dataCore.Email,
		Password:        dataCore.Password,
		Role:            dataCore.Role,
		Image:           dataCore.Image,
		AgencyCode:      dataCore.AgencyCode,
		NumberSchool:    dataCore.NumberSchool,
		AgreementLetter: dataCore.AgreementLetter,
		SuperAdminID:    dataCore.SuperAdminID,
	}
	return modelData
}
