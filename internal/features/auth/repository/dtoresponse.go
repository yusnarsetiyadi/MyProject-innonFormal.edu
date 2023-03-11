package repository

import "myproject/innonformaledu/internal/features/auth"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

func (data User) toCore() auth.UserCore {
	return auth.UserCore{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
