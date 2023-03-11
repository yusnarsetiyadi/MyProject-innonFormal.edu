package repository

import "myproject/innonformaledu/internal/features/schooladministrator"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

func (dataModel *SchoolAdministrator) toCore() schooladministrator.SchoolAdministratorCore {
	return schooladministrator.SchoolAdministratorCore{
		ID:              dataModel.ID,
		Name:            dataModel.Name,
		Email:           dataModel.Email,
		Password:        dataModel.Password,
		Role:            dataModel.Role,
		Image:           dataModel.Image,
		AgencyCode:      dataModel.AgencyCode,
		NumberSchool:    dataModel.NumberSchool,
		AgreementLetter: dataModel.AgreementLetter,
		SuperAdminID:    dataModel.SuperAdminID,
		CreatedAt:       dataModel.CreatedAt,
		UpdatedAt:       dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []SchoolAdministrator) []schooladministrator.SchoolAdministratorCore {
	var dataCore []schooladministrator.SchoolAdministratorCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
