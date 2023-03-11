package feedback

import "time"

/*

	Struct Core For Data Transfer Object

*/

type FeedbackCore struct {
	ID        uint
	Rating    float64
	Feedback  string
	StudentID uint
	ClassID   uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
