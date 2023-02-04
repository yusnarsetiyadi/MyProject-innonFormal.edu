package repository

import (
	"myproject/innonformaledu/features/feedback"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model
	Rating    float64
	Feedback  string
	StudentID uint
	ClassID   uint
}

func fromCore(dataCore feedback.FeedbackCore) Feedback {
	modelData := Feedback{
		Rating:    dataCore.Rating,
		Feedback:  dataCore.Feedback,
		StudentID: dataCore.StudentID,
		ClassID:   dataCore.ClassID,
	}
	return modelData
}

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
