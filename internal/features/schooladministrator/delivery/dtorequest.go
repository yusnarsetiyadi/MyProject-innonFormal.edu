package delivery

import "myproject/innonformaledu/internal/features/schooladministrator"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func toCore(input SchoolAdministratorRequest) schooladministrator.SchoolAdministratorCore {
	dataCore := schooladministrator.SchoolAdministratorCore{
		Name:            input.Name,
		Email:           input.Email,
		Password:        input.Password,
		Role:            input.Role,
		Image:           input.Image,
		AgencyCode:      input.AgencyCode,
		NumberSchool:    input.NumberSchool,
		AgreementLetter: input.AgreementLetter,
		SuperAdminID:    input.SuperAdminID,
	}
	return dataCore
}
