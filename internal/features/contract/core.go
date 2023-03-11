package contract

import "time"

/*

	Struct Core For Data Transfer Object

*/

type ContractCore struct {
	ID                    uint
	AgreementLetter       string
	Cost                  uint
	Portofolio            string
	ContractImage         string
	Duration              time.Time
	TeacherID             uint
	SchoolAdministratorID uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
