package delivery

import "myproject/innonformaledu/features/contract"

type ContractRequest struct {
	AgreementLetter       string `json:"agreement_letter" form:"agreement_letter"`
	Cost                  uint   `json:"cost" form:"cost"`
	Portofolio            string `json:"portofolio" form:"portofolio"`
	ContractImage         string `json:"contract_image" form:"contract_image"`
	TeacherID             uint   `json:"teacher_id" form:"teacher_id"`
	SchoolAdministratorID uint   `json:"school_administrator_id" form:"school_administrator_id"`
}

func toCore(input ContractRequest) contract.ContractCore {
	dataCore := contract.ContractCore{
		AgreementLetter:       input.AgreementLetter,
		Cost:                  input.Cost,
		Portofolio:            input.Portofolio,
		ContractImage:         input.ContractImage,
		TeacherID:             input.TeacherID,
		SchoolAdministratorID: input.SchoolAdministratorID,
	}
	return dataCore
}
