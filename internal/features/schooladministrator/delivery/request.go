package delivery

/*

	Struct Request

*/

type SchoolAdministratorRequest struct {
	Name            string `json:"name" form:"name"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	Role            string `json:"role" form:"role"`
	Image           string `json:"image" form:"image"`
	AgencyCode      uint   `json:"agency_code" form:"agency_code"`
	NumberSchool    uint   `json:"number_school" form:"number_school"`
	AgreementLetter string `json:"agreement_letter" form:"agreement_letter"`
	SuperAdminID    uint   `json:"super_admin_id" form:"super_admin_id"`
}
