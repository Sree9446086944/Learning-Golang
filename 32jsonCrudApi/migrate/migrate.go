package main

import (
	"github.com/Sree9446086944/jsonCrudApi/initializers"
	"github.com/Sree9446086944/jsonCrudApi/models"
)

//migrate our db, create table Post -> https://gorm.io/docs/index.html - refer

func init() {
	//load env and connect db before main
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema - migrate db, create table Post
	// db.AutoMigrate(&Product{})   - db - *gorm.DB
	initializers.DB.AutoMigrate(&models.Post{}) //DB var is available public in initializers, connect db, model Post get created in db in postgres

	// go run migrate/migrate.go -> this will run this file and migrate table in db

}
