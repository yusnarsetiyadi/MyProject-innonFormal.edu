package delivery

/*

	Struct Request

*/

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
