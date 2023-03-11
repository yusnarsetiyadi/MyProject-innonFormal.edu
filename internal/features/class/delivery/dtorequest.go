package delivery

import "myproject/innonformaledu/internal/features/class"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func toCore(input ClassRequest) class.ClassCore {
	dataCore := class.ClassCore{
		ClassName:             input.ClassName,
		ClassCategory:         input.ClassCategory,
		ClassDescription:      input.ClassDescription,
		ClassRule:             input.ClassRule,
		ClassImage:            input.ClassImage,
		AgencyCode:            input.AgencyCode,
		AverageRating:         input.AverageRating,
		SchoolAdministratorID: input.SchoolAdministratorID,
	}
	return dataCore
}
