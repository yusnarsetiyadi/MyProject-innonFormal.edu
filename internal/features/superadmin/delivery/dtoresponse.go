package delivery

import "myproject/innonformaledu/internal/features/superadmin"

/*

	Data Transfer Object
	Response Data From Struct Core

*/

func fromCore(dataCore superadmin.SuperAdminCore) SuperAdminResponse {
	return SuperAdminResponse{
		ID:        dataCore.ID,
		Name:      dataCore.Name,
		Email:     dataCore.Email,
		Password:  dataCore.Password,
		Role:      dataCore.Role,
		Image:     dataCore.Image,
		CreatedAt: dataCore.CreatedAt,
		UpdatedAt: dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []superadmin.SuperAdminCore) []SuperAdminResponse {
	var dataResponse []SuperAdminResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
