package repository

import "myproject/innonformaledu/internal/features/superadmin"

/*

	Data Transfer Object
	Response Data To Struct Core

*/

func (dataModel *SuperAdmin) toCore() superadmin.SuperAdminCore {
	return superadmin.SuperAdminCore{
		ID:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		Role:      dataModel.Role,
		Image:     dataModel.Image,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func toCoreList(dataModel []SuperAdmin) []superadmin.SuperAdminCore {
	var dataCore []superadmin.SuperAdminCore
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
