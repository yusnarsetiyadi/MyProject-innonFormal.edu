package delivery

import "myproject/innonformaledu/features/raport"

type RaportRequest struct {
	Mark      float64 `json:"mark" form:"mark"`
	Status    string  `json:"status" form:"status"`
	Comment   string  `json:"comment" form:"comment"`
	TeacherID uint    `json:"teacher_id" form:"teacher_id"`
}

func toCore(input RaportRequest) raport.RaportCore {
	dataCore := raport.RaportCore{
		Mark:      input.Mark,
		Status:    input.Status,
		Comment:   input.Comment,
		TeacherID: input.TeacherID,
	}
	return dataCore
}
