package delivery

import (
	"time"
)

/*

	Struct Response

*/

type SchoolAdministratorResponse struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	Role            string    `json:"role"`
	Image           string    `json:"image"`
	AgencyCode      uint      `json:"agency_code"`
	NumberSchool    uint      `json:"number_school"`
	AgreementLetter string    `json:"agreement_letter"`
	SuperAdminID    uint      `json:"super_admin_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
