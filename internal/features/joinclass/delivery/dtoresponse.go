package delivery

import "myproject/innonformaledu/internal/features/joinclass"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore joinclass.JoinClassCore) JoinClassResponse {
	return JoinClassResponse{
		ID:             dataCore.ID,
		JoinDate:       dataCore.JoinDate,
		PurposeFollow:  dataCore.PurposeFollow,
		InterestTalent: dataCore.InterestTalent,
		StudentID:      dataCore.StudentID,
		ClassID:        dataCore.ClassID,
		CreatedAt:      dataCore.CreatedAt,
		UpdatedAt:      dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []joinclass.JoinClassCore) []JoinClassResponse {
	var dataResponse []JoinClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
