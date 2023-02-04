package service

import (
	"myproject/innonformaledu/features/joinclass"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type joinclassService struct {
	joinclassRepository joinclass.RepositoryInterface
	validate            *validator.Validate
}

func New(repo joinclass.RepositoryInterface) joinclass.ServiceInterface {
	return &joinclassService{
		joinclassRepository: repo,
		validate:            validator.New(),
	}
}

func (service *joinclassService) GetAll() (data []joinclass.JoinClassCore, err error) {
	data, err = service.joinclassRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *joinclassService) GetById(id uint) (data joinclass.JoinClassCore, err error) {
	data, err = service.joinclassRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return joinclass.JoinClassCore{}, err
	}

	return data, nil
}

func (service *joinclassService) Create(input joinclass.JoinClassCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.joinclassRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *joinclassService) Update(input joinclass.JoinClassCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.joinclassRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *joinclassService) Delete(id uint) error {
	errDelete := service.joinclassRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
