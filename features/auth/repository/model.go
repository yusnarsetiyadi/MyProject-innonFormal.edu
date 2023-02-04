package repository

import (
	"myproject/innonformaledu/features/auth"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
	Role     string `validate:"required"`
}

func (data User) toCore() auth.UserCore {
	return auth.UserCore{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		Role:      data.Role,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
