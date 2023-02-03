package postgresql

import (
	"fmt"
	"log"
	"myproject/innonformaledu/config"

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

	// migrateDB(db)

	return db
}

// func migrateDB(db *gorm.DB){

// }
