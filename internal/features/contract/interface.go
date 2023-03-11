package contract

/*

	Interface Application

*/

type ServiceInterface interface {
	GetAll() (data []ContractCore, err error)
	Create(input ContractCore) error
	GetById(id uint) (data ContractCore, err error)
	Update(input ContractCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []ContractCore, err error)
	Create(input ContractCore) error
	GetById(id uint) (data ContractCore, err error)
	Update(input ContractCore, id uint) error
	Delete(id uint) error
}
