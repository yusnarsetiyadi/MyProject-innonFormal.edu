package delivery

import "myproject/innonformaledu/internal/features/student"

/*

	Data Transfer Object
	Request Data To Struct Core

*/

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
