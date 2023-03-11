package delivery

import (
	"time"
)

/*

	Struct Response

*/

type StudentResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Image      string    `json:"image"`
	AgencyCode uint      `json:"agency_code"`
	Gender     string    `json:"gender"`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
