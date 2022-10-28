package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	//from GORM doc , for postgres -> https://gorm.io/docs/connecting_to_the_database.html
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	//move this to .env file
	dsn := os.Getenv("DB_URL") //from godotenv package docs
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) //DB and err already declared externally so no :=, direct assign

	if err != nil {
		log.Fatal("Failed to connect to db.")
	}
}
