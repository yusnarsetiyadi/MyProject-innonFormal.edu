package delivery

import (
	"myproject/innonformaledu/features/auth"
)

type UserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToCore(userReq UserRequest) auth.UserCore {
	userCore := auth.UserCore{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
