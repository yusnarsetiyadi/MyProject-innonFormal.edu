package repository

import "myproject/innonformaledu/internal/features/feedback"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

func (dataModel *Feedback) toCore() feedback.FeedbackCore {
	return feedback.FeedbackCore{
		ID:        dataModel.ID,
		Rating:    dataModel.Rating,
		Feedback:  dataModel.Feedback,
		StudentID: dataModel.StudentID,
		ClassID:   dataModel.ClassID,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []Feedback) []feedback.FeedbackCore {
	var dataCore []feedback.FeedbackCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
