package factory

import (
	authDelivery "myproject/innonformaledu/internal/features/auth/delivery"
	authRepo "myproject/innonformaledu/internal/features/auth/repository"
	authService "myproject/innonformaledu/internal/features/auth/service"

	classDelivery "myproject/innonformaledu/internal/features/class/delivery"
	classRepo "myproject/innonformaledu/internal/features/class/repository"
	classService "myproject/innonformaledu/internal/features/class/service"

	contractDelivery "myproject/innonformaledu/internal/features/contract/delivery"
	contractRepo "myproject/innonformaledu/internal/features/contract/repository"
	contractService "myproject/innonformaledu/internal/features/contract/service"

	feedbackDelivery "myproject/innonformaledu/internal/features/feedback/delivery"
	feedbackRepo "myproject/innonformaledu/internal/features/feedback/repository"
	feedbackService "myproject/innonformaledu/internal/features/feedback/service"

	joinclassDelivery "myproject/innonformaledu/internal/features/joinclass/delivery"
	joinclassRepo "myproject/innonformaledu/internal/features/joinclass/repository"
	joinclassService "myproject/innonformaledu/internal/features/joinclass/service"

	raportDelivery "myproject/innonformaledu/internal/features/raport/delivery"
	raportRepo "myproject/innonformaledu/internal/features/raport/repository"
	raportService "myproject/innonformaledu/internal/features/raport/service"

	schooladministratorDelivery "myproject/innonformaledu/internal/features/schooladministrator/delivery"
	schooladministratorRepo "myproject/innonformaledu/internal/features/schooladministrator/repository"
	schooladministratorService "myproject/innonformaledu/internal/features/schooladministrator/service"

	studentDelivery "myproject/innonformaledu/internal/features/student/delivery"
	studentRepo "myproject/innonformaledu/internal/features/student/repository"
	studentService "myproject/innonformaledu/internal/features/student/service"

	teacherDelivery "myproject/innonformaledu/internal/features/teacher/delivery"
	teacherRepo "myproject/innonformaledu/internal/features/teacher/repository"
	teacherService "myproject/innonformaledu/internal/features/teacher/service"

	superadminDelivery "myproject/innonformaledu/internal/features/superadmin/delivery"
	superadminRepo "myproject/innonformaledu/internal/features/superadmin/repository"
	superadminService "myproject/innonformaledu/internal/features/superadmin/service"

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

	teacherRepoFactory := teacherRepo.New(db)
	teacherServiceFactory := teacherService.New(teacherRepoFactory)
	teacherDelivery.New(teacherServiceFactory, e)

	superadminRepoFactory := superadminRepo.New(db)
	superadminServiceFactory := superadminService.New(superadminRepoFactory)
	superadminDelivery.New(superadminServiceFactory, e)

}
