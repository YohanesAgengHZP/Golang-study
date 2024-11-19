package initializer

import (
	// "fmt"
	"log"
	// "os"
	// "path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnvVar() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}