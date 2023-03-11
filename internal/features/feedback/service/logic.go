package service

import (
	"myproject/innonformaledu/internal/features/feedback"
	"myproject/innonformaledu/utils/helper"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
)

/*

	Bussiness Logic

*/

type feedbackService struct {
	feedbackRepository feedback.RepositoryInterface
	validate           *validator.Validate
}

func New(repo feedback.RepositoryInterface) feedback.ServiceInterface {
	return &feedbackService{
		feedbackRepository: repo,
		validate:           validator.New(),
	}
}

func (service *feedbackService) GetAll() (data []feedback.FeedbackCore, err error) {
	data, err = service.feedbackRepository.GetAll()
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return nil, err
	}

	return data, nil
}

func (service *feedbackService) GetById(id uint) (data feedback.FeedbackCore, err error) {
	data, err = service.feedbackRepository.GetById(id)
	if err != nil {
		log.Error(err.Error())
		helper.LogDebug(err)
		return feedback.FeedbackCore{}, err
	}

	return data, nil
}

func (service *feedbackService) Create(input feedback.FeedbackCore) (err error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errCreate := service.feedbackRepository.Create(input)
	if errCreate != nil {
		log.Error(errCreate.Error())
		helper.LogDebug(errCreate)
		return errCreate
	}

	return nil
}

func (service *feedbackService) Update(input feedback.FeedbackCore, id uint) error {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}

	errUpdate := service.feedbackRepository.Update(input, id)
	if errUpdate != nil {
		log.Error(errUpdate.Error())
		helper.LogDebug(errUpdate)
		return errUpdate
	}

	return nil
}

func (service *feedbackService) Delete(id uint) error {
	errDelete := service.feedbackRepository.Delete(id)
	if errDelete != nil {
		log.Error(errDelete.Error())
		helper.LogDebug(errDelete)
		return errDelete
	}

	return nil
}
