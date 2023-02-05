package repository

import (
	"myproject/innonformaledu/features/schooladministrator"

	"gorm.io/gorm"
)

type schooladministratorRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) schooladministrator.RepositoryInterface {
	return &schooladministratorRepository{
		db: db,
	}
}

func (repo *schooladministratorRepository) GetAll() (data []schooladministrator.SchoolAdministratorCore, err error) {
	var results []SchoolAdministrator

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *schooladministratorRepository) GetById(id uint) (data schooladministrator.SchoolAdministratorCore, err error) {
	var result SchoolAdministrator

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

func (repo *schooladministratorRepository) Create(input schooladministrator.SchoolAdministratorCore) error {
	schooladministratorGorm := fromCore(input)

	tx := repo.db.Create(&schooladministratorGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *schooladministratorRepository) Update(input schooladministrator.SchoolAdministratorCore, id uint) error {
	schooladministratorGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&schooladministratorGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *schooladministratorRepository) Delete(id uint) error {
	var schooladministrator SchoolAdministrator

	tx := repo.db.Delete(&schooladministrator, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
