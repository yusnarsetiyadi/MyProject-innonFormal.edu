package mysql

import (
	"fmt"
	"log"
	"myproject/innonformaledu/config"

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

	// migrateDB(db)

	return db
}

// func migrateDB(db *gorm.DB){

// }