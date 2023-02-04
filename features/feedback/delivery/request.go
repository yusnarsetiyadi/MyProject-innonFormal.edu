package delivery

import "myproject/innonformaledu/features/feedback"

type FeedbackRequest struct {
	Rating    float64 `json:"rating" form:"rating"`
	Feedback  string  `json:"feedback" form:"feedback"`
	StudentID uint    `json:"student_id" form:"student_id"`
	ClassID   uint    `json:"class_id" form:"class_id"`
}

func toCore(input FeedbackRequest) feedback.FeedbackCore {
	dataCore := feedback.FeedbackCore{
		Rating:    input.Rating,
		Feedback:  input.Feedback,
		StudentID: input.StudentID,
		ClassID:   input.ClassID,
	}
	return dataCore
}
