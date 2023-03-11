package student

import "time"

/*

	Struct Core For Data Transfer Object

*/

type StudentCore struct {
	ID         uint
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
