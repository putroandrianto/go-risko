package connections

import (
	"fmt"
	"go-risko/config"
	"go-risko/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Config(DB_HOST), config.Config(DB_PORT), config.Config(DB_USER), config.Config(DB_PASSWORD), config.Config(DB_NAME))
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error when connecting to database")
	} else {
		log.Println("Connected to database")
	}

	DBMigrate()
}

func DBMigrate() {
	DB.AutoMigrate(&models.User{}, &models.RiskProfile{})
	fmt.Println("Database migrated")
}