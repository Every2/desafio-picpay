package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig(fileName string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(fileName)
	config.AddConfigPath(".")
	config.AddConfigPath("$HOME")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("Error while parser configuration file", err)
	}

	return config
}