package main

import (
	"log"

	"github.com/Every2/desafio-picpay/config"
	"github.com/Every2/desafio-picpay/server"
)

func main() {
	log.Println("Starting app...")
	config := config.InitConfig("secrets.toml")
	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	log.Printf("Temporary use of database: %v", dbHandler)
}