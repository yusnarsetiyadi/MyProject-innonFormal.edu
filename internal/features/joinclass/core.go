package joinclass

import "time"

/*

	Struct Core For Data Transfer Object

*/

type JoinClassCore struct {
	ID             uint
	JoinDate       time.Time
	PurposeFollow  string
	InterestTalent string
	ClassID        uint
	StudentID      uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
