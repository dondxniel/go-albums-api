package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var error error
	
	// database connection string
	dsn := os.Getenv("DB_CONNECTION_STRING");

	// create DB connection
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{});

	// panic if error exists
	if error != nil {
		panic("Database connection failed");
	}

	DB.AutoMigrate()
	
}