package main

// This File migration for automation Migration
// For example !!!!!
// Dont Touch it!!!!
// Use Database Synctax

import (
	"go-server/database/connection"
	"go-server/database/models"
	"go-server/initializes"
)

func init() {
	initializes.LoadVariable()
	connection.ConnectToDb()
}

func main() {
	connection.DB.AutoMigrate(&models.User{})
}
