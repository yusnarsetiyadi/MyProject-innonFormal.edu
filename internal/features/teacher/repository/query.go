package repository

import (
	"myproject/innonformaledu/internal/features/teacher"

	"gorm.io/gorm"
)

/*

	Bussiness Query

*/

type teacherRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) teacher.RepositoryInterface {
	return &teacherRepository{
		db: db,
	}
}

func (repo *teacherRepository) GetAll() (data []teacher.TeacherCore, err error) {
	var results []Teacher

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *teacherRepository) GetById(id uint) (data teacher.TeacherCore, err error) {
	var result Teacher

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

func (repo *teacherRepository) Create(input teacher.TeacherCore) error {
	teacherGorm := fromCore(input)

	tx := repo.db.Create(&teacherGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *teacherRepository) Update(input teacher.TeacherCore, id uint) error {
	teacherGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&teacherGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *teacherRepository) Delete(id uint) error {
	var teacher Teacher

	tx := repo.db.Delete(&teacher, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
