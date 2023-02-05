package delivery

import (
	"myproject/innonformaledu/features/student"
	"time"
)

type StudentResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Image      string    `json:"image"`
	AgencyCode uint      `json:"agency_code"`
	Gender     string    `json:"gender"`
	Address    string    `json:"address"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func fromCore(dataCore student.StudentCore) StudentResponse {
	return StudentResponse{
		ID:         dataCore.ID,
		Name:       dataCore.Name,
		Email:      dataCore.Email,
		Password:   dataCore.Password,
		Role:       dataCore.Role,
		Image:      dataCore.Image,
		AgencyCode: dataCore.AgencyCode,
		Gender:     dataCore.Gender,
		Address:    dataCore.Address,
		CreatedAt:  dataCore.CreatedAt,
		UpdatedAt:  dataCore.UpdatedAt,
	}
}

func fromCoreList(dataCore []student.StudentCore) []StudentResponse {
	var dataResponse []StudentResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, fromCore(v))
	}
	return dataResponse
}
