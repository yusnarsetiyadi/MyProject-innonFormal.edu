package delivery

import "myproject/innonformaledu/internal/features/feedback"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func toCore(input FeedbackRequest) feedback.FeedbackCore {
	dataCore := feedback.FeedbackCore{
		Rating:    input.Rating,
		Feedback:  input.Feedback,
		StudentID: input.StudentID,
		ClassID:   input.ClassID,
	}
	return dataCore
}
