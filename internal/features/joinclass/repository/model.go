package repository

import (
	"time"

	"gorm.io/gorm"
)

/*

	Struct Model For Database

*/

type JoinClass struct {
	gorm.Model
	JoinDate       time.Time
	PurposeFollow  string
	InterestTalent string
	ClassID        uint
	StudentID      uint
}
