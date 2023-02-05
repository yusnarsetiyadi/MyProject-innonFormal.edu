package delivery

import (
	"myproject/innonformaledu/features/superadmin"
	"time"
)

type SuperAdminResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

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
