package repository

import (
	"myproject/innonformaledu/features/feedback"

	"gorm.io/gorm"
)

type feedbackRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) feedback.RepositoryInterface {
	return &feedbackRepository{
		db: db,
	}
}

func (repo *feedbackRepository) GetAll() (data []feedback.FeedbackCore, err error) {
	var results []Feedback

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *feedbackRepository) GetById(id uint) (data feedback.FeedbackCore, err error) {
	var result Feedback

	tx := repo.db.First(&result, id)
	if tx.Error != nil {
		return data, tx.Error
	}

	if tx.RowsAffected == 0 {
		return data, tx.Error
	}

	var dataCore = result.toCore()
	return dataCore, nil
}

func (repo *feedbackRepository) Create(input feedback.FeedbackCore) error {
	feedbackGorm := fromCore(input)

	tx := repo.db.Create(&feedbackGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *feedbackRepository) Update(input feedback.FeedbackCore, id uint) error {
	feedbackGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&feedbackGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *feedbackRepository) Delete(id uint) error {
	var feedback Feedback

	tx := repo.db.Delete(&feedback, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
