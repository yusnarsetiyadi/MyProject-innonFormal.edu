package delivery

import "myproject/innonformaledu/features/teacher"

type TeacherRequest struct {
	Name       string `json:"name" form:"name"`
	Email      string `json:"email" form:"email"`
	Password   string `json:"password" form:"password"`
	Role       string `json:"role" form:"role"`
	Image      string `json:"image" form:"image"`
	AgencyCode uint   `json:"agency_code" form:"agency_code"`
	Gender     string `json:"gender" form:"gender"`
	Address    string `json:"address" form:"address"`
	Portofolio string `json:"portofolio" form:"portofolio"`
	Expertise  string `json:"expertise" form:"expertise"`
	Material   string `json:"material" form:"material"`
}

func toCore(input TeacherRequest) teacher.TeacherCore {
	dataCore := teacher.TeacherCore{
		Name:       input.Name,
		Email:      input.Email,
		Password:   input.Password,
		Role:       input.Role,
		Image:      input.Image,
		AgencyCode: input.AgencyCode,
		Gender:     input.Gender,
		Address:    input.Address,
		Portofolio: input.Portofolio,
		Expertise:  input.Expertise,
		Material:   input.Material,
	}
	return dataCore
}
