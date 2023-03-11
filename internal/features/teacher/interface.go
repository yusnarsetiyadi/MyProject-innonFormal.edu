package teacher

/*

	Interface Application

*/

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
