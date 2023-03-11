package delivery

import "myproject/innonformaledu/internal/features/raport"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func toCore(input RaportRequest) raport.RaportCore {
	dataCore := raport.RaportCore{
		Mark:      input.Mark,
		Status:    input.Status,
		Comment:   input.Comment,
		TeacherID: input.TeacherID,
	}
	return dataCore
}
