package postgresql

import (
	"fmt"
	"log"
	"myproject/innonformaledu/config"

	class "myproject/innonformaledu/features/class"
	contract "myproject/innonformaledu/features/contract"
	feedback "myproject/innonformaledu/features/feedback"
	joinclass "myproject/innonformaledu/features/joinclass"
	raport "myproject/innonformaledu/features/raport"
	schooladministrator "myproject/innonformaledu/features/schooladministrator"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgreSQLDB(cfg *config.AppConfig) *gorm.DB {
	postgresqlDB := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", cfg.DB_HOST, cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT)
	db, err := gorm.Open(postgres.Open(postgresqlDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	fmt.Print("connection postgres success")

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
}
