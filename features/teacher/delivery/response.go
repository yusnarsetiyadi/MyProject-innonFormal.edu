package delivery

import (
	"myproject/innonformaledu/features/teacher"
	"time"
)

type TeacherResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Image      string    `json:"image"`
	AgencyCode uint      `json:"agency_code"`
	Gender     string    `json:"gender"`
	Address    string    `json:"address"`
	Portofolio string    `json:"portofolio"`
	Expertise  string    `json:"expertise"`
	Material   string    `json:"material"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func fromCore(dataCore teacher.TeacherCore) TeacherResponse {
	return TeacherResponse{
		ID:         dataCore.ID,
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		Role:       dataCore.Role,
		Image:      dataCore.Image,
		AgencyCode: dataCore.AgencyCode,
		Gender:     dataCore.Gender,
		Address:    dataCore.Address,
		Portofolio: dataCore.Portofolio,
		Expertise:  dataCore.Expertise,
		Material:   dataCore.Material,
		CreatedAt:  dataCore.CreatedAt,
		UpdatedAt:  dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []teacher.TeacherCore) []TeacherResponse {
	var dataResponse []TeacherResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
