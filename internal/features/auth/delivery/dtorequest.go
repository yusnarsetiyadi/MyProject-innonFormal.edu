package delivery

import "myproject/innonformaledu/internal/features/auth"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func ToCore(userReq UserRequest) auth.UserCore {
	userCore := auth.UserCore{
		Email:    userReq.Email,
		Password: userReq.Password,
	}
	return userCore
}
