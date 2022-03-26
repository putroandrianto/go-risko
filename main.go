package main

import (
	"go-risko/app"
	"go-risko/db"
)

func main() {
	db.ConnectDB()
	app.StartApp()
}