package repository

import (
	"time"

	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type Contract struct {
	gorm.Model
	AgreementLetter       string
	Cost                  uint
	Portofolio            string
	ContractImage         string
	Duration              time.Time
	TeacherID             uint
	SchoolAdministratorID uint
}
