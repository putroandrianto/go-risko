package models

import (
	"go-risko/connections"
	"go-risko/data"

	"gorm.io/gorm"
)

func CreateUser(user data.User) *gorm.DB {
	result := connections.DB.Create(&user)
	return result
}

func GetUsers(limit int, offset int) []data.User {
	users := []data.User{}
	connections.DB.
		Select("id, name, age").
		Limit(limit).
		Offset(offset).
		Find(&users)
	return users
}

func GetUserById(userId string) (data.User, *gorm.DB) {
	var user data.User
	result := connections.DB.Select("users.id, users.name, users.age, RiskProfile.*").Joins("RiskProfile").Where("users.id = ?", userId).First(&user)
	return user, result
}