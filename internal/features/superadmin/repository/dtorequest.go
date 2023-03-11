package repository

import "myproject/innonformaledu/internal/features/superadmin"

/*

	Data Transfer Object
	Request Data From Struct Core

*/

func fromCore(dataCore superadmin.SuperAdminCore) SuperAdmin {
	modelData := SuperAdmin{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
		Role:     dataCore.Role,
		Image:    dataCore.Image,
	}
	return modelData
}
