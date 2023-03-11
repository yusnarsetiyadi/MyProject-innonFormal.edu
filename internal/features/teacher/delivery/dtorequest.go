package delivery

import "myproject/innonformaledu/internal/features/teacher"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

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
