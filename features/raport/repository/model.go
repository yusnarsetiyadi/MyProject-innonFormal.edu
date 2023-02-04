package repository

import (
	"myproject/innonformaledu/features/raport"

	"gorm.io/gorm"
)

type Raport struct {
	gorm.Model
	StudentID uint
	Mark      float64
	Status    string
	Comment   string
	TeacherID uint
}

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
