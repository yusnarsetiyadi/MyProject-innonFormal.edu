package delivery

import (
	"myproject/innonformaledu/features/joinclass"
	"time"
)

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

func fromCore(dataCore joinclass.JoinClassCore) JoinClassResponse {
	return JoinClassResponse{
		ID:             dataCore.ID,
		JoinDate:       dataCore.JoinDate,
		PurposeFollow:  dataCore.PurposeFollow,
		InterestTalent: dataCore.InterestTalent,
		StudentID:      dataCore.StudentID,
		ClassID:        dataCore.ClassID,
		CreatedAt:      dataCore.CreatedAt,
		UpdatedAt:      dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []joinclass.JoinClassCore) []JoinClassResponse {
	var dataResponse []JoinClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
