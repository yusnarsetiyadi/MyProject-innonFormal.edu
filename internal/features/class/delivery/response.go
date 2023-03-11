package delivery

import (
	"time"
)

/*

	Struct Response

*/

type ClassResponse struct {
	ID                    uint      `json:"id"`
	ClassName             string    `json:"class_name"`
	ClassCategory         string    `json:"class_category"`
	ClassDescription      string    `json:"class_description"`
	ClassRule             string    `json:"class_rule"`
	ClassImage            string    `json:"class_image"`
	AgencyCode            uint      `json:"agency_code"`
	AverageRating         float64   `json:"average_rating"`
	SchoolAdministratorID uint      `json:"school_administrator_id"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
