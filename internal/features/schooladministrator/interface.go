package schooladministrator

/*

	Interface Application

*/

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
