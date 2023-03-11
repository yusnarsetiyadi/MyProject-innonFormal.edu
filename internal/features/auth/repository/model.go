package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}
