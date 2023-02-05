package service

import (
	"myproject/innonformaledu/features/teacher"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

type teacherService struct {
	teacherRepository teacher.RepositoryInterface
	validate          *validator.Validate
}

func New(repo teacher.RepositoryInterface) teacher.ServiceInterface {
	return &teacherService{
		teacherRepository: repo,
		validate:          validator.New(),
	}
}

func (service *teacherService) GetAll() (data []teacher.TeacherCore, err error) {
	data, err = service.teacherRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *teacherService) GetById(id uint) (data teacher.TeacherCore, err error) {
	data, err = service.teacherRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return teacher.TeacherCore{}, err
	}

	return data, nil
}

func (service *teacherService) Create(input teacher.TeacherCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.teacherRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *teacherService) Update(input teacher.TeacherCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.teacherRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *teacherService) Delete(id uint) error {
	errDelete := service.teacherRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
