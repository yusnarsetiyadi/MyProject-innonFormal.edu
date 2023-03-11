package repository

import (
	"myproject/innonformaledu/internal/features/joinclass"

	"gorm.io/gorm"
)

/*

	Bussiness Query

*/

type joinclassRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) joinclass.RepositoryInterface {
	return &joinclassRepository{
		db: db,
	}
}

func (repo *joinclassRepository) GetAll() (data []joinclass.JoinClassCore, err error) {
	var results []JoinClass

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *joinclassRepository) GetById(id uint) (data joinclass.JoinClassCore, err error) {
	var result JoinClass

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

func (repo *joinclassRepository) Create(input joinclass.JoinClassCore) error {
	joinclassGorm := fromCore(input)

	tx := repo.db.Create(&joinclassGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *joinclassRepository) Update(input joinclass.JoinClassCore, id uint) error {
	joinclassGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&joinclassGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *joinclassRepository) Delete(id uint) error {
	var joinclass JoinClass

	tx := repo.db.Delete(&joinclass, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
