package repository

import "myproject/innonformaledu/internal/features/raport"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

func (dataModel *Raport) toCore() raport.RaportCore {
	return raport.RaportCore{
		StudentID: dataModel.StudentID,
		Mark:      dataModel.Mark,
		Status:    dataModel.Status,
		Comment:   dataModel.Comment,
		TeacherID: dataModel.TeacherID,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []Raport) []raport.RaportCore {
	var dataCore []raport.RaportCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
