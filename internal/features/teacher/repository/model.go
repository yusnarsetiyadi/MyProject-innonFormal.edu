package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type Teacher struct {
	gorm.Model
	ClassID    uint
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
	Portofolio string
	Expertise  string
	Material   string
}
