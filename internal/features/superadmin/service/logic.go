package service

import (
	"myproject/innonformaledu/internal/features/superadmin"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

/*

	Bussiness Logic

*/

type superadminService struct {
	superadminRepository superadmin.RepositoryInterface
	validate             *validator.Validate
}

func New(repo superadmin.RepositoryInterface) superadmin.ServiceInterface {
	return &superadminService{
		superadminRepository: repo,
		validate:             validator.New(),
	}
}

func (service *superadminService) GetAll() (data []superadmin.SuperAdminCore, err error) {
	data, err = service.superadminRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *superadminService) GetById(id uint) (data superadmin.SuperAdminCore, err error) {
	data, err = service.superadminRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return superadmin.SuperAdminCore{}, err
	}

	return data, nil
}

func (service *superadminService) Create(input superadmin.SuperAdminCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.superadminRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *superadminService) Update(input superadmin.SuperAdminCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.superadminRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *superadminService) Delete(id uint) error {
	errDelete := service.superadminRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
