package service

import (
	"errors"
	"fmt"
	"myproject/innonformaledu/features/auth"
	"myproject/innonformaledu/middlewares"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	authData auth.RepositoryInterface
	validate *validator.Validate
}

func New(data auth.RepositoryInterface) auth.ServiceInterface {
	return &authService{
		authData: data,
		validate: validator.New(),
	}
}

func (service *authService) Login(dataCore auth.UserCore) (data auth.UserCore, token string, err error) {

	if errValidate := service.validate.Struct(dataCore); errValidate != nil {
		log.Error(errValidate.Error())
		return auth.UserCore{}, "", errors.New("Failed to login, error validate input, please check your input.")
	}

	result, errLogin := service.authData.FindUser(dataCore.Email)
	if errLogin != nil {
		log.Error(errLogin.Error())
		if strings.Contains(errLogin.Error(), "table") {
			return auth.UserCore{}, "", errors.New("Failed to login, error on request, please contact your administrator.")
		} else if strings.Contains(errLogin.Error(), "found") {
			return auth.UserCore{}, "", errors.New("Failed to login, email not found, please check email again.")
		} else {
			return auth.UserCore{}, "", errors.New("Failed to login, other error, please contact your administrator.")
		}
	}

	errCheckPass := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(dataCore.Password))
	fmt.Println("\n\n Data Core User = ", dataCore)
	fmt.Println("\n\n Result User= ", result)
	if errCheckPass != nil {
		log.Error(errCheckPass.Error())
		return auth.UserCore{}, "", errors.New("Failed to login, password didn't match, please check password again.")
	}

	token, errToken := middlewares.CreateToken(uint(result.ID), result.Role)
	if errToken != nil {
		log.Error(errToken.Error())
		return auth.UserCore{}, "", errors.New("Failed to login, error on generate token, please try again.")
	}

	return result, token, nil
}
