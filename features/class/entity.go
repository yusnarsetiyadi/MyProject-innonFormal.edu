package class

import "time"

type ClassCore struct {
	ID                    uint
	ClassName             string
	ClassCategory         string
	ClassDescription      string
	ClassRule             string
	ClassImage            string
	AgencyCode            uint
	AverageRating         float64
	SchoolAdministratorID uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

type ServiceInterface interface {
	GetAll() (data []ClassCore, err error)
	Create(input ClassCore) error
	GetById(id uint) (data ClassCore, err error)
	Update(input ClassCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []ClassCore, err error)
	Create(input ClassCore) error
	GetById(id uint) (data ClassCore, err error)
	Update(input ClassCore, id uint) error
	Delete(id uint) error
}
