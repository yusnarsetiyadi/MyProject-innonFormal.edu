package repository

import (
	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type Feedback struct {
	gorm.Model
	Rating    float64
	Feedback  string
	StudentID uint
	ClassID   uint
}
