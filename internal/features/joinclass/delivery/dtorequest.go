package delivery

import "myproject/innonformaledu/internal/features/joinclass"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func toCore(input JoinClassRequest) joinclass.JoinClassCore {
	dataCore := joinclass.JoinClassCore{
		JoinDate:       input.JoinDate,
		PurposeFollow:  input.PurposeFollow,
		InterestTalent: input.InterestTalent,
		StudentID:      input.StudentID,
		ClassID:        input.ClassID,
	}
	return dataCore
}
