package superadmin

/*

	Interface Application

*/

type ServiceInterface interface {
	GetAll() (data []SuperAdminCore, err error)
	Create(input SuperAdminCore) error
	GetById(id uint) (data SuperAdminCore, err error)
	Update(input SuperAdminCore, id uint) error
	Delete(id uint) error
}

type RepositoryInterface interface {
	GetAll() (data []SuperAdminCore, err error)
	Create(input SuperAdminCore) error
	GetById(id uint) (data SuperAdminCore, err error)
	Update(input SuperAdminCore, id uint) error
	Delete(id uint) error
}
