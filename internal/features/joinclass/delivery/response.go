package delivery

import (
	"time"
)

/*

	Struct Response

*/

type JoinClassResponse struct {
	ID             uint      `json:"id"`
	JoinDate       time.Time `json:"join_date"`
	PurposeFollow  string    `json:"purpose_follow"`
	InterestTalent string    `json:"interest_talent"`
	StudentID      uint      `json:"student_id"`
	ClassID        uint      `json:"class_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
