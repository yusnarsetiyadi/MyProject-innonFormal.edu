package delivery

import "myproject/innonformaledu/internal/features/auth"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func FromCore(dataCore auth.UserCore, token string) UserResponse {
	return UserResponse{
		ID:    dataCore.ID,
		Name:  dataCore.Name,
		Email: dataCore.Email,
		Role:  dataCore.Role,
		Token: token,
	}
}
