package delivery

/*

	Struct Request

*/

type FeedbackRequest struct {
	Rating    float64 `json:"rating" form:"rating"`
	Feedback  string  `json:"feedback" form:"feedback"`
	StudentID uint    `json:"student_id" form:"student_id"`
	ClassID   uint    `json:"class_id" form:"class_id"`
}
