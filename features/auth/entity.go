package auth

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	Login(input UserCore) (data UserCore, token string, err error)
}

type RepositoryInterface interface {
	FindUser(email string) (result UserCore, err error)
}
