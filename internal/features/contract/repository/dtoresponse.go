package repository

import "myproject/innonformaledu/internal/features/contract"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

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
