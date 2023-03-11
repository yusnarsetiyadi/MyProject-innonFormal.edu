package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type Raport struct {
	gorm.Model
	StudentID uint
	Mark      float64
	Status    string
	Comment   string
	TeacherID uint
}
