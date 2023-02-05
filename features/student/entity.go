package student

import "time"

type StudentCore struct {
	ID         uint
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceInterface interface {
	GetAll() (data []StudentCore, err error)
	Create(input StudentCore) error
	GetById(id uint) (data StudentCore, err error)
	Update(input StudentCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []StudentCore, err error)
	Create(input StudentCore) error
	GetById(id uint) (data StudentCore, err error)
	Update(input StudentCore, id uint) error
	Delete(id uint) error
}
