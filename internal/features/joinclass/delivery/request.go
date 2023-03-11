package delivery

import (
	"time"
)

/*

	Struct Request

*/

type JoinClassRequest struct {
	JoinDate       time.Time `json:"join_date" form:"join_date"`
	PurposeFollow  string    `json:"purpose_follow" form:"purpose_follow"`
	InterestTalent string    `json:"interest_talent" form:"interest_talent"`
	StudentID      uint      `json:"student_id" form:"student_id"`
	ClassID        uint      `json:"class_id" form:"class_id"`
}
