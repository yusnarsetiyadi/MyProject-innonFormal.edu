package delivery

import (
	"time"
)

/*

	Struct Response

*/

type RaportResponse struct {
	StudentID uint      `json:"student_id"`
	Mark      float64   `json:"mark"`
	Status    string    `json:"status"`
	Comment   string    `json:"comment"`
	TeacherID uint      `json:"teacher_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
