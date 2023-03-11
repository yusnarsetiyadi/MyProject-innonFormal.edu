package raport

/*

	Interface Application

*/

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
