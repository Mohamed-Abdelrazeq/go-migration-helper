package helpers

import (
	"fmt"
	"os"
)

func ScanDatabaseInfo() (string, string) {
	var driver string
	var databaseString string

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
	Cache("GMH_DB_DRIVER", driver)
	Cache("GMH_DB_STRING", databaseString)

	return driver, databaseString
}
