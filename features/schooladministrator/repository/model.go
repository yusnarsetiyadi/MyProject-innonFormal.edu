package repository

import (
	"myproject/innonformaledu/features/schooladministrator"

	"gorm.io/gorm"
)

type SchoolAdministrator struct {
	gorm.Model
	Name            string
	Email           string
	Password        string
	Role            string
	Image           string
	AgencyCode      uint
	NumberSchool    uint
	AgreementLetter string
	SuperAdminID    uint
}

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
