package superadmin

import "time"

type SuperAdminCore struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Role      string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll() (data []SuperAdminCore, err error)
	Create(input SuperAdminCore) error
	GetById(id uint) (data SuperAdminCore, err error)
	Update(input SuperAdminCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []SuperAdminCore, err error)
	Create(input SuperAdminCore) error
	GetById(id uint) (data SuperAdminCore, err error)
	Update(input SuperAdminCore, id uint) error
	Delete(id uint) error
}
