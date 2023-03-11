package repository

import "myproject/innonformaledu/internal/features/joinclass"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

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
