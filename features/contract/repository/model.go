package repository

import (
	"myproject/innonformaledu/features/contract"
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	AgreementLetter       string
	Cost                  uint
	Portofolio            string
	ContractImage         string
	Duration              time.Time
	TeacherID             uint
	SchoolAdministratorID uint
}

func fromCore(dataCore contract.ContractCore) Contract {
	modelData := Contract{
		AgreementLetter:       dataCore.AgreementLetter,
		Cost:                  dataCore.Cost,
		Portofolio:            dataCore.Portofolio,
		ContractImage:         dataCore.ContractImage,
		TeacherID:             dataCore.TeacherID,
		SchoolAdministratorID: dataCore.SchoolAdministratorID,
	}
	return modelData
}

func (dataModel *Contract) toCore() contract.ContractCore {
	return contract.ContractCore{
		ID:                    dataModel.ID,
		AgreementLetter:       dataModel.AgreementLetter,
		Cost:                  dataModel.Cost,
		Portofolio:            dataModel.Portofolio,
		ContractImage:         dataModel.ContractImage,
		TeacherID:             dataModel.TeacherID,
		SchoolAdministratorID: dataModel.SchoolAdministratorID,
		CreatedAt:             dataModel.CreatedAt,
		UpdatedAt:             dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []Contract) []contract.ContractCore {
	var dataCore []contract.ContractCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
