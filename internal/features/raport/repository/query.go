package repository

import (
	"myproject/innonformaledu/internal/features/raport"

	"gorm.io/gorm"
)

/*

	Bussiness Query

*/

type raportRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) raport.RepositoryInterface {
	return &raportRepository{
		db: db,
	}
}

func (repo *raportRepository) GetAll() (data []raport.RaportCore, err error) {
	var results []Raport

	tx := repo.db.Find(&results)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataCore = toCoreList(results)
	return dataCore, nil
}

func (repo *raportRepository) GetById(id uint) (data raport.RaportCore, err error) {
	var result Raport

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

func (repo *raportRepository) Create(input raport.RaportCore) error {
	raportGorm := fromCore(input)

	tx := repo.db.Create(&raportGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *raportRepository) Update(input raport.RaportCore, id uint) error {
	raportGorm := fromCore(input)

	tx := repo.db.Where("id = ?", id).Updates(&raportGorm)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}

func (repo *raportRepository) Delete(id uint) error {
	var raport Raport

	tx := repo.db.Delete(&raport, id)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	return nil
}
