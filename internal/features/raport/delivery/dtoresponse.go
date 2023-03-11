package delivery

import "myproject/innonformaledu/internal/features/raport"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore raport.RaportCore) RaportResponse {
	return RaportResponse{
		StudentID: dataCore.StudentID,
		Mark:      dataCore.Mark,
		Status:    dataCore.Status,
		Comment:   dataCore.Comment,
		TeacherID: dataCore.TeacherID,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []raport.RaportCore) []RaportResponse {
	var dataResponse []RaportResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
