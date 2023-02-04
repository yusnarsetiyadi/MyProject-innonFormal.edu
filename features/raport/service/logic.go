package service

import (
	"myproject/innonformaledu/features/raport"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type raportService struct {
	raportRepository raport.RepositoryInterface
	validate         *validator.Validate
}

func New(repo raport.RepositoryInterface) raport.ServiceInterface {
	return &raportService{
		raportRepository: repo,
		validate:         validator.New(),
	}
}

func (service *raportService) GetAll() (data []raport.RaportCore, err error) {
	data, err = service.raportRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *raportService) GetById(id uint) (data raport.RaportCore, err error) {
	data, err = service.raportRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return raport.RaportCore{}, err
	}

	return data, nil
}

func (service *raportService) Create(input raport.RaportCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.raportRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *raportService) Update(input raport.RaportCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.raportRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *raportService) Delete(id uint) error {
	errDelete := service.raportRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
