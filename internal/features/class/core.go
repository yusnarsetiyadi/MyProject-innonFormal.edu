package class

import "time"

/*

	Struct Core For Data Transfer Object

*/

type ClassCore struct {
	ID                    uint
	ClassName             string
	ClassCategory         string
	ClassDescription      string
	ClassRule             string
	ClassImage            string
	AgencyCode            uint
	AverageRating         float64
	SchoolAdministratorID uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
