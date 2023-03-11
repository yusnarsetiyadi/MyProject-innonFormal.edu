package delivery

import "myproject/innonformaledu/internal/features/feedback"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

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
