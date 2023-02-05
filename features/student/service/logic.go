package service

import (
	"myproject/innonformaledu/features/student"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type studentService struct {
	studentRepository student.RepositoryInterface
	validate          *validator.Validate
}

func New(repo student.RepositoryInterface) student.ServiceInterface {
	return &studentService{
		studentRepository: repo,
		validate:          validator.New(),
	}
}

func (service *studentService) GetAll() (data []student.StudentCore, err error) {
	data, err = service.studentRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *studentService) GetById(id uint) (data student.StudentCore, err error) {
	data, err = service.studentRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return student.StudentCore{}, err
	}

	return data, nil
}

func (service *studentService) Create(input student.StudentCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.studentRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *studentService) Update(input student.StudentCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.studentRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *studentService) Delete(id uint) error {
	errDelete := service.studentRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
