package student

/*

	Interface Application

*/

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
