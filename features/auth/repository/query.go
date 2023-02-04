package repository

import (
	"errors"
	"myproject/innonformaledu/features/auth"

	"gorm.io/gorm"
)

type authData struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.RepositoryInterface {
	return &authData{
		db: db,
	}
}

func (repo *authData) FindUser(email string) (result auth.UserCore, err error) {
	var userData User
	tx := repo.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return auth.UserCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return auth.UserCore{}, errors.New("login failed")
	}

	result = userData.toCore()

	return result, nil
}
