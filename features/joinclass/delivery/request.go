package delivery

import (
	"myproject/innonformaledu/features/joinclass"
	"time"
)

type JoinClassRequest struct {
	JoinDate       time.Time `json:"join_date" form:"join_date"`
	PurposeFollow  string    `json:"purpose_follow" form:"purpose_follow"`
	InterestTalent string    `json:"interest_talent" form:"interest_talent"`
	StudentID      uint      `json:"student_id" form:"student_id"`
	ClassID        uint      `json:"class_id" form:"class_id"`
}

func toCore(input JoinClassRequest) joinclass.JoinClassCore {
	dataCore := joinclass.JoinClassCore{
		JoinDate:       input.JoinDate,
		PurposeFollow:  input.PurposeFollow,
		InterestTalent: input.InterestTalent,
		StudentID:      input.StudentID,
		ClassID:        input.ClassID,
	}
	return dataCore
}
