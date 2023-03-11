package delivery

import (
	"time"
)

/*

	Struct Request

*/

type ContractRequest struct {
	AgreementLetter       string    `json:"agreement_letter" form:"agreement_letter"`
	Cost                  uint      `json:"cost" form:"cost"`
	Portofolio            string    `json:"portofolio" form:"portofolio"`
	ContractImage         string    `json:"contract_image" form:"contract_image"`
	Duration              time.Time `json:"duration" form:"duration"`
	TeacherID             uint      `json:"teacher_id" form:"teacher_id"`
	SchoolAdministratorID uint      `json:"school_administrator_id" form:"school_administrator_id"`
}
