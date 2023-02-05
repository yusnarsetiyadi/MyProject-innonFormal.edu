package service

import (
	"myproject/innonformaledu/features/schooladministrator"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type schooladministratorService struct {
	schooladministratorRepository schooladministrator.RepositoryInterface
	validate                      *validator.Validate
}

func New(repo schooladministrator.RepositoryInterface) schooladministrator.ServiceInterface {
	return &schooladministratorService{
		schooladministratorRepository: repo,
		validate:                      validator.New(),
	}
}

func (service *schooladministratorService) GetAll() (data []schooladministrator.SchoolAdministratorCore, err error) {
	data, err = service.schooladministratorRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *schooladministratorService) GetById(id uint) (data schooladministrator.SchoolAdministratorCore, err error) {
	data, err = service.schooladministratorRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return schooladministrator.SchoolAdministratorCore{}, err
	}

	return data, nil
}

func (service *schooladministratorService) Create(input schooladministrator.SchoolAdministratorCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.schooladministratorRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *schooladministratorService) Update(input schooladministrator.SchoolAdministratorCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.schooladministratorRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *schooladministratorService) Delete(id uint) error {
	errDelete := service.schooladministratorRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
