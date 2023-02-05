package schooladministrator

import "time"

type SchoolAdministratorCore struct {
	ID              uint
	Name            string
	Email           string
	Password        string
	Role            string
	Image           string
	AgencyCode      uint
	NumberSchool    uint
	AgreementLetter string
	SuperAdminID    uint
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ServiceInterface interface {
	GetAll() (data []SchoolAdministratorCore, err error)
	Create(input SchoolAdministratorCore) error
	GetById(id uint) (data SchoolAdministratorCore, err error)
	Update(input SchoolAdministratorCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []SchoolAdministratorCore, err error)
	Create(input SchoolAdministratorCore) error
	GetById(id uint) (data SchoolAdministratorCore, err error)
	Update(input SchoolAdministratorCore, id uint) error
	Delete(id uint) error
}
