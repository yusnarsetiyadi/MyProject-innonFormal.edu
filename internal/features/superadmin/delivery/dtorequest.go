package delivery

import "myproject/innonformaledu/internal/features/superadmin"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

func toCore(input SuperAdminRequest) superadmin.SuperAdminCore {
	dataCore := superadmin.SuperAdminCore{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     input.Role,
		Image:    input.Image,
	}
	return dataCore
}
