package repository

import "myproject/innonformaledu/internal/features/class"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

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
