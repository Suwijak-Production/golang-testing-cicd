package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	ServerAddress string
}

var Config ConfigStruct

func Load() {

	err := godotenv.Load(filepath.Join(os.Getenv("APP_PATH"), ".env"))
	log.Printf("con Loading .env file from: %s", os.Getenv("APP_PATH"))
	if err != nil {
		log.Fatalf("Error loading .env file config")
	}

	Config = ConfigStruct{
		ServerAddress: os.Getenv("SERVER_ADDRESS"),
	}
}
