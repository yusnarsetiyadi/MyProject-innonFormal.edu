package repository

import (
	"myproject/innonformaledu/internal/features/student"

	"gorm.io/gorm"
)

/*

	Bussiness Query

*/

type studentRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) student.RepositoryInterface {
	return &studentRepository{
		db: db,
	}
}

func (repo *studentRepository) GetAll() (data []student.StudentCore, err error) {
	var results []Student

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *studentRepository) GetById(id uint) (data student.StudentCore, err error) {
	var result Student

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

func (repo *studentRepository) Create(input student.StudentCore) error {
	studentGorm := fromCore(input)

	tx := repo.db.Create(&studentGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *studentRepository) Update(input student.StudentCore, id uint) error {
	studentGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&studentGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *studentRepository) Delete(id uint) error {
	var student Student

	tx := repo.db.Delete(&student, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
