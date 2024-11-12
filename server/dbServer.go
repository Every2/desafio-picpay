package server

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func InitDatabase(config *viper.Viper) *sql.DB {
	connectionString := config.GetString("database.connection_string")
	driverName := config.GetString("database.driver_name")

	if connectionString == "" {
		log.Fatalf("Database connection string is empty")
	}

	dbHandler, err := sql.Open(driverName, connectionString)
	if err != nil {
		log.Fatalf("Error while initializing database: %v", err)
	}

	err = dbHandler.Ping()
	if err != nil {
		dbHandler.Close()
		log.Fatalf("Error while validating database: %v", err)
	}

	return dbHandler
}