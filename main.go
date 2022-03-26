package main

import (
	"go-risko/app"
	"go-risko/connections"
)

func main() {
	connections.ConnectDB()
	app.StartApp()
}