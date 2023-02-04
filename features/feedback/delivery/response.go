package delivery

import (
	"myproject/innonformaledu/features/feedback"
	"time"
)

type FeedbackResponse struct {
	ID        uint      `json:"id"`
	Rating    float64   `json:"rating"`
	Feedback  string    `json:"feedback"`
	StudentID uint      `json:"student_id"`
	ClassID   uint      `json:"class_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func fromCore(dataCore feedback.FeedbackCore) FeedbackResponse {
	return FeedbackResponse{
		ID:        dataCore.ID,
		Rating:    dataCore.Rating,
		Feedback:  dataCore.Feedback,
		StudentID: dataCore.StudentID,
		ClassID:   dataCore.ClassID,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []feedback.FeedbackCore) []FeedbackResponse {
	var dataResponse []FeedbackResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
