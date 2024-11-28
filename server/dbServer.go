package server

import (
	"log"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(config *viper.Viper) *gorm.DB {
	connectionString := config.GetString("database.connection_string")

	if connectionString == "" {
		log.Fatalf("Database connection string is empty")
	}

	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while initializing database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error while getting database handle: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Error while validating database connection: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}