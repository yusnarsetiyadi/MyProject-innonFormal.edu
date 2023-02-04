package contract

import "time"

type ContractCore struct {
	ID                    uint
	AgreementLetter       string
	Cost                  uint
	Portofolio            string
	ContractImage         string
	TeacherID             uint
	SchoolAdministratorID uint
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

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
