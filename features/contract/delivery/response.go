package delivery

import (
	"myproject/innonformaledu/features/contract"
	"time"
)

type ContractResponse struct {
	ID                    uint      `json:"id"`
	AgreementLetter       string    `json:"agreement_letter"`
	Cost                  uint      `json:"cost"`
	Portofolio            string    `json:"portofolio"`
	ContractImage         string    `json:"contract_image"`
	TeacherID             uint      `json:"teacher_id"`
	SchoolAdministratorID uint      `json:"school_administrator_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}

func fromCore(dataCore contract.ContractCore) ContractResponse {
	return ContractResponse{
		ID:                    dataCore.ID,
		AgreementLetter:       dataCore.AgreementLetter,
		Cost:                  dataCore.Cost,
		Portofolio:            dataCore.Portofolio,
		ContractImage:         dataCore.ContractImage,
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
