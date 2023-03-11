package repository

import "myproject/innonformaledu/internal/features/raport"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

func fromCore(dataCore raport.RaportCore) Raport {
	modelData := Raport{
		StudentID: dataCore.StudentID,
		Mark:      dataCore.Mark,
		Status:    dataCore.Status,
		Comment:   dataCore.Comment,
		TeacherID: dataCore.TeacherID,
	}
	return modelData
}
