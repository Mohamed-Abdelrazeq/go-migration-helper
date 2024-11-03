package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Cache(key, value string) {
	envFile, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Error opening .env file: ", err)
	}
	defer envFile.Close()

	_, err = envFile.WriteString(fmt.Sprintf("%s=%s\n", key, value))
	if err != nil {
		log.Fatal("Error writing to .env file: ", err)
	}
}

func InitCache() {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		file, err := os.Create(".env")
		if err != nil {
			log.Fatal("Error creating .env file: ", err)
		}
		defer file.Close()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
}
