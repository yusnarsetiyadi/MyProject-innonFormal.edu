package teacher

import "time"

type TeacherCore struct {
	ID         uint
	Name       string
	Email      string
	Password   string
	Role       string
	Image      string
	AgencyCode uint
	Gender     string
	Address    string
	Portofolio string
	Expertise  string
	Material   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ServiceInterface interface {
	GetAll() (data []TeacherCore, err error)
	Create(input TeacherCore) error
	GetById(id uint) (data TeacherCore, err error)
	Update(input TeacherCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []TeacherCore, err error)
	Create(input TeacherCore) error
	GetById(id uint) (data TeacherCore, err error)
	Update(input TeacherCore, id uint) error
	Delete(id uint) error
}
