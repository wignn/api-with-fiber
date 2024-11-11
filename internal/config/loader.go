package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Get() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("errror when loading config")
	}

	return &Config{
		Server: Server{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		Database: Database{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			Pass: os.Getenv("DB_PASS"),
			Name: os.Getenv("DB_NAME"),
			Tz:   os.Getenv("DB_TZ"),
		},
	}

}
