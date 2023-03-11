package delivery

import (
	"time"
)

/*

	Struct Response

*/

type FeedbackResponse struct {
	ID        uint      `json:"id"`
	Rating    float64   `json:"rating"`
	Feedback  string    `json:"feedback"`
	StudentID uint      `json:"student_id"`
	ClassID   uint      `json:"class_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
