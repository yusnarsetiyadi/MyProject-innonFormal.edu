package delivery

import "myproject/innonformaledu/features/class"

type ClassRequest struct {
	ClassName             string  `json:"class_name" form:"class_name"`
	ClassCategory         string  `json:"class_category" form:"class_category"`
	ClassDescription      string  `json:"class_description" form:"class_description"`
	ClassRule             string  `json:"class_rule" form:"class_rule"`
	ClassImage            string  `json:"class_image" form:"class_image"`
	AgencyCode            uint    `json:"agency_code" form:"agency_code"`
	AverageRating         float64 `json:"average_rating" form:"average_rating"`
	SchoolAdministratorID uint    `json:"school_administrator_id" form:"school_administrator_id"`
}

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
