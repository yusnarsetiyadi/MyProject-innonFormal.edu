package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type SchoolAdministrator struct {
	gorm.Model
	Name            string
	Email           string
	Password        string
	Role            string
	Image           string
	AgencyCode      uint
	NumberSchool    uint
	AgreementLetter string
	SuperAdminID    uint
}
