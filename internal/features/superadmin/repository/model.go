package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type SuperAdmin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
	Image    string
}
