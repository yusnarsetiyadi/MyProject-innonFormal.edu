package repository

import "myproject/innonformaledu/internal/features/class"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

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
