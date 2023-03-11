package mysql

import (
	"fmt"
	"log"
	"myproject/innonformaledu/internal/config"

	class "myproject/innonformaledu/internal/features/class"
	contract "myproject/innonformaledu/internal/features/contract"
	feedback "myproject/innonformaledu/internal/features/feedback"
	joinclass "myproject/innonformaledu/internal/features/joinclass"
	raport "myproject/innonformaledu/internal/features/raport"
	schooladministrator "myproject/innonformaledu/internal/features/schooladministrator"
	student "myproject/innonformaledu/internal/features/student"
	superadmin "myproject/innonformaledu/internal/features/superadmin"
	teacher "myproject/innonformaledu/internal/features/teacher"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySQLDB(cfg *config.AppConfig) *gorm.DB {
	mysqlDB := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	db, err := gorm.Open(mysql.Open(mysqlDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	fmt.Print("connection mysql success")

	if err != nil {
		log.Fatal("Cannot connect to DB")
	}

	MigrateDB(db)

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&class.ClassCore{})
	db.AutoMigrate(&contract.ContractCore{})
	db.AutoMigrate(&feedback.FeedbackCore{})
	db.AutoMigrate(&joinclass.JoinClassCore{})
	db.AutoMigrate(&raport.RaportCore{})
	db.AutoMigrate(&schooladministrator.SchoolAdministratorCore{})
	db.AutoMigrate(&student.StudentCore{})
	db.AutoMigrate(&teacher.TeacherCore{})
	db.AutoMigrate(&superadmin.SuperAdminCore{})
}
