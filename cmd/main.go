package main

import (
	"app/cmd/config"
	"app/db"
	"app/model"
	"log"
)

var database *db.DataStore
var configuration config.Configuration

func init() {
	// setup Configuration
	configuration = config.Load("dev")

	// setup DB connection
	database = db.NewDatabase(configuration)
}

func main() {
	// Initiate UP SQL Migrations
	// If fail will execute down migrations then exit the application
	// db.Migration(&configuration, false)
	err := database.Db.AutoMigrate(model.UserProfile{}, model.User{},
		model.Application{}, model.ReleaseTicket{}, model.Debt{})

	if err != nil {
		log.Fatal(err.Error())
	}
	// Setup Gin Route
	r := setupRoutes()
	r.Run(":8078")
}
