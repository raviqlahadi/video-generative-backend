package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var PythonServerUrl string

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	PythonServerUrl = os.Getenv("PYTHON_SERVER_URL")
	if PythonServerUrl == "" {
		log.Fatal("Error python server url not found")
	}
}
