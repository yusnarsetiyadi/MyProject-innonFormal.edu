package factory

import (
	authDelivery "myproject/innonformaledu/features/auth/delivery"
	authRepo "myproject/innonformaledu/features/auth/repository"
	authService "myproject/innonformaledu/features/auth/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)
}
