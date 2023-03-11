package delivery

import "myproject/innonformaledu/internal/features/contract"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func toCore(input ContractRequest) contract.ContractCore {
	dataCore := contract.ContractCore{
		AgreementLetter:       input.AgreementLetter,
		Cost:                  input.Cost,
		Portofolio:            input.Portofolio,
		ContractImage:         input.ContractImage,
		Duration:              input.Duration,
		TeacherID:             input.TeacherID,
		SchoolAdministratorID: input.SchoolAdministratorID,
	}
	return dataCore
}
