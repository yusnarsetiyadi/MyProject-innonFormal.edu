package repository

import (
	"myproject/innonformaledu/internal/features/class"

	"gorm.io/gorm"
)

/*

	Bussiness Query

*/

type classRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) class.RepositoryInterface {
	return &classRepository{
		db: db,
	}
}

func (repo *classRepository) GetAll() (data []class.ClassCore, err error) {
	var results []Class

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *classRepository) GetById(id uint) (data class.ClassCore, err error) {
	var result Class

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

func (repo *classRepository) Create(input class.ClassCore) error {
	classGorm := fromCore(input)

	tx := repo.db.Create(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *classRepository) Update(input class.ClassCore, id uint) error {
	classGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&classGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *classRepository) Delete(id uint) error {
	var class Class

	tx := repo.db.Delete(&class, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
