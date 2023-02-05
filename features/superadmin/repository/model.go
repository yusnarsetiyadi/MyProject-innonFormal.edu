package repository

import (
	"myproject/innonformaledu/features/superadmin"

	"gorm.io/gorm"
)

type SuperAdmin struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
	Image    string
}

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
