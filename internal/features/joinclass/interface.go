package joinclass

/*

	Interface Application

*/

type ServiceInterface interface {
	GetAll() (data []JoinClassCore, err error)
	Create(input JoinClassCore) error
	GetById(id uint) (data JoinClassCore, err error)
	Update(input JoinClassCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []JoinClassCore, err error)
	Create(input JoinClassCore) error
	GetById(id uint) (data JoinClassCore, err error)
	Update(input JoinClassCore, id uint) error
	Delete(id uint) error
}
