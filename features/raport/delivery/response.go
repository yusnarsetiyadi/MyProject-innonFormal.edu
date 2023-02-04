package delivery

import (
	"myproject/innonformaledu/features/raport"
	"time"
)

type RaportResponse struct {
	StudentID uint      `json:"student_id"`
	Mark      float64   `json:"mark"`
	Status    string    `json:"status"`
	Comment   string    `json:"comment"`
	TeacherID uint      `json:"teacher_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func fromCore(dataCore raport.RaportCore) RaportResponse {
	return RaportResponse{
		StudentID: dataCore.StudentID,
		Mark:      dataCore.Mark,
		Status:    dataCore.Status,
		Comment:   dataCore.Comment,
		TeacherID: dataCore.TeacherID,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []raport.RaportCore) []RaportResponse {
	var dataResponse []RaportResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
