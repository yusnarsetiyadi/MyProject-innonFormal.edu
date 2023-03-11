package raport

import "time"

/*

	Struct Core For Data Transfer Object

*/

type RaportCore struct {
	StudentID uint
	Mark      float64
	Status    string
	Comment   string
	TeacherID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
