package delivery

/*

	Struct Request

*/

type RaportRequest struct {
	Mark      float64 `json:"mark" form:"mark"`
	Status    string  `json:"status" form:"status"`
	Comment   string  `json:"comment" form:"comment"`
	TeacherID uint    `json:"teacher_id" form:"teacher_id"`
}
