package delivery

import "myproject/innonformaledu/internal/features/contract"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore contract.ContractCore) ContractResponse {
	return ContractResponse{
		ID:                    dataCore.ID,
		AgreementLetter:       dataCore.AgreementLetter,
		Cost:                  dataCore.Cost,
		Portofolio:            dataCore.Portofolio,
		ContractImage:         dataCore.ContractImage,
		Duration:              dataCore.Duration,
		TeacherID:             dataCore.TeacherID,
		SchoolAdministratorID: dataCore.SchoolAdministratorID,
		CreatedAt:             dataCore.CreatedAt,
		UpdatedAt:             dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []contract.ContractCore) []ContractResponse {
	var dataResponse []ContractResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
