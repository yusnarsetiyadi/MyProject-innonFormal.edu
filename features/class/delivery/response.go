package delivery

import (
	"myproject/innonformaledu/features/class"
	"time"
)

type ClassResponse struct {
	ID                    uint      `json:"id" form:"id"`
	ClassName             string    `json:"class_name" form:"class_name"`
	ClassCategory         string    `json:"class_category" form:"class_category"`
	ClassDescription      string    `json:"class_description" form:"class_description"`
	ClassRule             string    `json:"class_rule" form:"class_rule"`
	ClassImage            string    `json:"class_image" form:"class_image"`
	AgencyCode            uint      `json:"agency_code" form:"agency_code"`
	AverageRating         float64   `json:"average_rating" form:"average_rating"`
	SchoolAdministratorID uint      `json:"school_administrator_id" form:"school_administrator_id"`
	CreatedAt             time.Time `json:"created_at" form:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" form:"updated_at"`
}

func fromCore(dataCore class.ClassCore) ClassResponse {
	return ClassResponse{
		ID:                    dataCore.ID,
		ClassName:             dataCore.ClassName,
		ClassCategory:         dataCore.ClassCategory,
		ClassDescription:      dataCore.ClassDescription,
		ClassRule:             dataCore.ClassRule,
		ClassImage:            dataCore.ClassImage,
		AgencyCode:            dataCore.AgencyCode,
		AverageRating:         dataCore.AverageRating,
		SchoolAdministratorID: dataCore.SchoolAdministratorID,
		CreatedAt:             dataCore.CreatedAt,
		UpdatedAt:             dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []class.ClassCore) []ClassResponse {
	var dataResponse []ClassResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}