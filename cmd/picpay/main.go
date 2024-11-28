package main

import (
	"fmt"
	"log"

	"github.com/Every2/desafio-picpay/config"
	"github.com/Every2/desafio-picpay/models"
	"github.com/Every2/desafio-picpay/server"
)

func main() {
	log.Println("Starting app...")
	config := config.InitConfig("example.toml")
	log.Println("Initializing database")
	db := server.InitDatabase(config)
	err := db.AutoMigrate(&models.User{}, &models.Transaction{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}
	httpServer := server.InitHttpServer(config, db)
	httpServer.Start()

	fmt.Println("App is running...")
}