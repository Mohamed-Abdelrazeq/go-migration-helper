package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ScanDatabaseInfo() (string, string) {
	var driver string
	var databaseString string

	// Create .env file if it doesn't exist
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

	// Check if the database info is cached
	if os.Getenv("GMH_DB_DRIVER") != "" && os.Getenv("GMH_DB_STRING") != "" {
		return os.Getenv("GMH_DB_DRIVER"), os.Getenv("GMH_DB_STRING")
	}

	// Scan the database info
	fmt.Print("Enter database driver: ")
	fmt.Scanln(&driver)
	fmt.Print("Enter database string: ")
	fmt.Scanln(&databaseString)

	// Cache the database info
	envFile, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Error opening .env file: ", err)
	}
	defer envFile.Close()

	_, err = envFile.WriteString(fmt.Sprintf("GMH_DB_DRIVER=%s\nGMH_DB_STRING=%s\n", driver, databaseString))
	if err != nil {
		log.Fatal("Error writing to .env file: ", err)
	}

	return driver, databaseString
}
