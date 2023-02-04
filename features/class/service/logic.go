package service

import (
	"myproject/innonformaledu/features/class"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type classService struct {
	classRepository class.RepositoryInterface
	validate        *validator.Validate
}

func New(repo class.RepositoryInterface) class.ServiceInterface {
	return &classService{
		classRepository: repo,
		validate:        validator.New(),
	}
}

func (service *classService) GetAll() (data []class.ClassCore, err error) {
	data, err = service.classRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *classService) GetById(id uint) (data class.ClassCore, err error) {
	data, err = service.classRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return class.ClassCore{}, err
	}

	return data, nil
}

func (service *classService) Create(input class.ClassCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.classRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *classService) Update(input class.ClassCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.classRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *classService) Delete(id uint) error {
	errDelete := service.classRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
