package delivery

import (
	"myproject/innonformaledu/features/schooladministrator"
	"time"
)

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

func fromCore(dataCore schooladministrator.SchoolAdministratorCore) SchoolAdministratorResponse {
	return SchoolAdministratorResponse{
		ID:              dataCore.ID,
		Name:            dataCore.Name,
		Email:           dataCore.Email,
		Password:        dataCore.Password,
		Role:            dataCore.Role,
		Image:           dataCore.Image,
		AgencyCode:      dataCore.AgencyCode,
		NumberSchool:    dataCore.NumberSchool,
		AgreementLetter: dataCore.AgreementLetter,
		SuperAdminID:    dataCore.SuperAdminID,
		CreatedAt:       dataCore.CreatedAt,
		UpdatedAt:       dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []schooladministrator.SchoolAdministratorCore) []SchoolAdministratorResponse {
	var dataResponse []SchoolAdministratorResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
