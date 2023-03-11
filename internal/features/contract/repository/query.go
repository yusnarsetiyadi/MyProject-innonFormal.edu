package repository

import (
	"myproject/innonformaledu/internal/features/contract"

	"gorm.io/gorm"
)

/*

	Bussiness Query

*/

type contractRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) contract.RepositoryInterface {
	return &contractRepository{
		db: db,
	}
}

func (repo *contractRepository) GetAll() (data []contract.ContractCore, err error) {
	var results []Contract

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *contractRepository) GetById(id uint) (data contract.ContractCore, err error) {
	var result Contract

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

func (repo *contractRepository) Create(input contract.ContractCore) error {
	contractGorm := fromCore(input)

	tx := repo.db.Create(&contractGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *contractRepository) Update(input contract.ContractCore, id uint) error {
	contractGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&contractGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *contractRepository) Delete(id uint) error {
	var contract Contract

	tx := repo.db.Delete(&contract, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
