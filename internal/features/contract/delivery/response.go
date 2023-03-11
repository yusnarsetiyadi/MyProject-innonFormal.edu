package delivery

import (
	"time"
)

/*

	Struct Response

*/

type ContractResponse struct {
	ID                    uint      `json:"id"`
	AgreementLetter       string    `json:"agreement_letter"`
	Cost                  uint      `json:"cost"`
	Portofolio            string    `json:"portofolio"`
	ContractImage         string    `json:"contract_image"`
	Duration              time.Time `json:"duration"`
	TeacherID             uint      `json:"teacher_id"`
	SchoolAdministratorID uint      `json:"school_administrator_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
