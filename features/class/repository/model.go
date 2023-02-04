package repository

import (
	"myproject/innonformaledu/features/class"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName             string
	ClassCategory         string
	ClassDescription      string
	ClassRule             string
	ClassImage            string
	AgencyCode            uint
	AverageRating         float64
	SchoolAdministratorID uint
}

func fromCore(dataCore class.ClassCore) Class {
	modelData := Class{
		ClassName:             dataCore.ClassName,
		ClassCategory:         dataCore.ClassCategory,
		ClassDescription:      dataCore.ClassDescription,
		ClassRule:             dataCore.ClassRule,
		ClassImage:            dataCore.ClassImage,
		AgencyCode:            dataCore.AgencyCode,
		AverageRating:         dataCore.AverageRating,
		SchoolAdministratorID: dataCore.SchoolAdministratorID,
	}
	return modelData
}

func (dataModel *Class) toCore() class.ClassCore {
	return class.ClassCore{
		ID:                    dataModel.ID,
		ClassName:             dataModel.ClassName,
		ClassCategory:         dataModel.ClassCategory,
		ClassDescription:      dataModel.ClassDescription,
		ClassRule:             dataModel.ClassRule,
		ClassImage:            dataModel.ClassImage,
		AgencyCode:            dataModel.AgencyCode,
		AverageRating:         dataModel.AverageRating,
		SchoolAdministratorID: dataModel.SchoolAdministratorID,
		CreatedAt:             dataModel.CreatedAt,
		UpdatedAt:             dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []Class) []class.ClassCore {
	var dataCore []class.ClassCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
