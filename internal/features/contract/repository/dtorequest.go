package repository

import "myproject/innonformaledu/internal/features/contract"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

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
