package connections

import (
	"fmt"
	"go-risko/config"
	"go-risko/data"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDB() {
	dsn := fmt.Sprintf("%v:%v@/%v?charset=utf8&parseTime=True", config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println("Error when connecting to database")
	} else {
		log.Println("Connected to database")
	}

	DBMigrate()
}

func DBMigrate() {
	DB.AutoMigrate(&data.User{}, &data.RiskProfile{})
	fmt.Println("Database migrated")
}

// func DBConnect() {
// 	database.DBConn, err = gorm.Open("mysql", "go-risko.db")
// 	if err != nil {
// 		panic("failed to connect database")
// 	}
// 	fmt.Println("Connected to database")
// 	database.DBConn.AutoMigrate(&models.User{}, &models.RiskProfile{})
// 	fmt.Println("Database migrated")
// }