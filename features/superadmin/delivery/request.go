package delivery

import "myproject/innonformaledu/features/superadmin"

type SuperAdminRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
	Image    string `json:"image" form:"image"`
}

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
