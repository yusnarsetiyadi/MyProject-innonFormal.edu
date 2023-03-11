package schooladministrator

import "time"

/*

	Struct Core For Data Transfer Object

*/

type SchoolAdministratorCore struct {
	ID              uint
	Name            string
	Email           string
	Password        string
	Role            string
	Image           string
	AgencyCode      uint
	NumberSchool    uint
	AgreementLetter string
	SuperAdminID    uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
