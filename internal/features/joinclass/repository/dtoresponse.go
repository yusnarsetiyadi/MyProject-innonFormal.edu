package repository

import "myproject/innonformaledu/internal/features/joinclass"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

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
