package service

import (
	"myproject/innonformaledu/internal/features/contract"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

/*

	Bussiness Logic

*/

type contractService struct {
	contractRepository contract.RepositoryInterface
	validate           *validator.Validate
}

func New(repo contract.RepositoryInterface) contract.ServiceInterface {
	return &contractService{
		contractRepository: repo,
		validate:           validator.New(),
	}
}

func (service *contractService) GetAll() (data []contract.ContractCore, err error) {
	data, err = service.contractRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *contractService) GetById(id uint) (data contract.ContractCore, err error) {
	data, err = service.contractRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return contract.ContractCore{}, err
	}

	return data, nil
}

func (service *contractService) Create(input contract.ContractCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.contractRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *contractService) Update(input contract.ContractCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.contractRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *contractService) Delete(id uint) error {
	errDelete := service.contractRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
