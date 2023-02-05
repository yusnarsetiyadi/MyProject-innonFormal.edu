package repository

import (
	"myproject/innonformaledu/features/superadmin"

	"gorm.io/gorm"
)

type superadminRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) superadmin.RepositoryInterface {
	return &superadminRepository{
		db: db,
	}
}

func (repo *superadminRepository) GetAll() (data []superadmin.SuperAdminCore, err error) {
	var results []SuperAdmin

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *superadminRepository) GetById(id uint) (data superadmin.SuperAdminCore, err error) {
	var result SuperAdmin

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

func (repo *superadminRepository) Create(input superadmin.SuperAdminCore) error {
	superadminGorm := fromCore(input)

	tx := repo.db.Create(&superadminGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *superadminRepository) Update(input superadmin.SuperAdminCore, id uint) error {
	superadminGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&superadminGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *superadminRepository) Delete(id uint) error {
	var superadmin SuperAdmin

	tx := repo.db.Delete(&superadmin, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
