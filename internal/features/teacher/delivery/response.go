package delivery

import (
	"time"
)

/*

	Struct Response

*/

type TeacherResponse struct {
	ClassID    uint      `json:"class_id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Image      string    `json:"image"`
	AgencyCode uint      `json:"agency_code"`
	Gender     string    `json:"gender"`
	Address    string    `json:"address"`
	Portofolio string    `json:"portofolio"`
	Expertise  string    `json:"expertise"`
	Material   string    `json:"material"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
