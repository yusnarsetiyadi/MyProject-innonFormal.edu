package repository

import (
	"myproject/innonformaledu/features/joinclass"
	"time"

	"gorm.io/gorm"
)

type JoinClass struct {
	gorm.Model
	JoinDate       time.Time
	PurposeFollow  string
	InterestTalent string
	ClassID        uint
	StudentID      uint
}

func fromCore(dataCore joinclass.JoinClassCore) JoinClass {
	modelData := JoinClass{
		JoinDate:       dataCore.JoinDate,
		PurposeFollow:  dataCore.PurposeFollow,
		InterestTalent: dataCore.InterestTalent,
		StudentID:      dataCore.StudentID,
		ClassID:        dataCore.ClassID,
	}
	return modelData
}

func (dataModel *JoinClass) toCore() joinclass.JoinClassCore {
	return joinclass.JoinClassCore{
		ID:             dataModel.ID,
		JoinDate:       dataModel.JoinDate,
		PurposeFollow:  dataModel.PurposeFollow,
		InterestTalent: dataModel.InterestTalent,
		StudentID:      dataModel.StudentID,
		ClassID:        dataModel.ClassID,
		CreatedAt:      dataModel.CreatedAt,
		UpdatedAt:      dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []JoinClass) []joinclass.JoinClassCore {
	var dataCore []joinclass.JoinClassCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
