package factory

import (
	authDelivery "myproject/innonformaledu/features/auth/delivery"
	authRepo "myproject/innonformaledu/features/auth/repository"
	authService "myproject/innonformaledu/features/auth/service"

	classDelivery "myproject/innonformaledu/features/class/delivery"
	classRepo "myproject/innonformaledu/features/class/repository"
	classService "myproject/innonformaledu/features/class/service"

	contractDelivery "myproject/innonformaledu/features/contract/delivery"
	contractRepo "myproject/innonformaledu/features/contract/repository"
	contractService "myproject/innonformaledu/features/contract/service"

	feedbackDelivery "myproject/innonformaledu/features/feedback/delivery"
	feedbackRepo "myproject/innonformaledu/features/feedback/repository"
	feedbackService "myproject/innonformaledu/features/feedback/service"

	joinclassDelivery "myproject/innonformaledu/features/joinclass/delivery"
	joinclassRepo "myproject/innonformaledu/features/joinclass/repository"
	joinclassService "myproject/innonformaledu/features/joinclass/service"

	raportDelivery "myproject/innonformaledu/features/raport/delivery"
	raportRepo "myproject/innonformaledu/features/raport/repository"
	raportService "myproject/innonformaledu/features/raport/service"

	schooladministratorDelivery "myproject/innonformaledu/features/schooladministrator/delivery"
	schooladministratorRepo "myproject/innonformaledu/features/schooladministrator/repository"
	schooladministratorService "myproject/innonformaledu/features/schooladministrator/service"

	studentDelivery "myproject/innonformaledu/features/student/delivery"
	studentRepo "myproject/innonformaledu/features/student/repository"
	studentService "myproject/innonformaledu/features/student/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {

	authRepoFactory := authRepo.New(db)
	authServiceFactory := authService.New(authRepoFactory)
	authDelivery.New(authServiceFactory, e)

	classRepoFactory := classRepo.New(db)
	classServiceFactory := classService.New(classRepoFactory)
	classDelivery.New(classServiceFactory, e)

	contractRepoFactory := contractRepo.New(db)
	contractServiceFactory := contractService.New(contractRepoFactory)
	contractDelivery.New(contractServiceFactory, e)

	feedbackRepoFactory := feedbackRepo.New(db)
	feedbackServiceFactory := feedbackService.New(feedbackRepoFactory)
	feedbackDelivery.New(feedbackServiceFactory, e)

	joinclassRepoFactory := joinclassRepo.New(db)
	joinclassServiceFactory := joinclassService.New(joinclassRepoFactory)
	joinclassDelivery.New(joinclassServiceFactory, e)

	raportRepoFactory := raportRepo.New(db)
	raportServiceFactory := raportService.New(raportRepoFactory)
	raportDelivery.New(raportServiceFactory, e)

	schooladministratorRepoFactory := schooladministratorRepo.New(db)
	schooladministratorServiceFactory := schooladministratorService.New(schooladministratorRepoFactory)
	schooladministratorDelivery.New(schooladministratorServiceFactory, e)

	studentRepoFactory := studentRepo.New(db)
	studentServiceFactory := studentService.New(studentRepoFactory)
	studentDelivery.New(studentServiceFactory, e)

}
