package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type Student struct {
	gorm.Model
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
}
