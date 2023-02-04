package auth

import "time"

type UserCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
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
