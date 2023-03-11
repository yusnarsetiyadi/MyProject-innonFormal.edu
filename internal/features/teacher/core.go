package teacher

import "time"

/*

	Struct Core For Data Transfer Object

*/

type TeacherCore struct {
	ClassID    uint
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
	Portofolio string
	Expertise  string
	Material   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
