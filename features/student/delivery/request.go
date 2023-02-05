package delivery

import "myproject/innonformaledu/features/student"

type StudentRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Role       string `json:"role" form:"role"`
	Image      string `json:"image" form:"image"`
	AgencyCode uint   `json:"agency_code" form:"agency_code"`
	Gender     string `json:"gender" form:"gender"`
	Address    string `json:"address" form:"address"`
}

func toCore(input StudentRequest) student.StudentCore {
	dataCore := student.StudentCore{
		Name:       input.Name,
		Email:      input.Email,
		Password:   input.Password,
		Role:       input.Role,
		Image:      input.Image,
		AgencyCode: input.AgencyCode,
		Gender:     input.Gender,
		Address:    input.Address,
	}
	return dataCore
}
