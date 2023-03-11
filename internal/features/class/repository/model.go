package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type Class struct {
	gorm.Model
	ClassName             string
	ClassCategory         string
	ClassDescription      string
	ClassRule             string
	ClassImage            string
	AgencyCode            uint
	AverageRating         float64
	SchoolAdministratorID uint
}
