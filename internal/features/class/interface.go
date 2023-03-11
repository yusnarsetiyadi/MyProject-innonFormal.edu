package class

/*

	Interface Application

*/

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
