package repository

import "myproject/innonformaledu/internal/features/feedback"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

func fromCore(dataCore feedback.FeedbackCore) Feedback {
	modelData := Feedback{
		Rating:    dataCore.Rating,
		Feedback:  dataCore.Feedback,
		StudentID: dataCore.StudentID,
		ClassID:   dataCore.ClassID,
	}
	return modelData
}
