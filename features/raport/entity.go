package raport

import "time"

type RaportCore struct {
	StudentID uint
	Mark      float64
	Status    string
	Comment   string
	TeacherID uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll() (data []RaportCore, err error)
	Create(input RaportCore) error
	GetById(id uint) (data RaportCore, err error)
	Update(input RaportCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []RaportCore, err error)
	Create(input RaportCore) error
	GetById(id uint) (data RaportCore, err error)
	Update(input RaportCore, id uint) error
	Delete(id uint) error
}
