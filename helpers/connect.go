package helpers

import (
	"database/sql"
	"log"
)

func Connect(driver string, databaseString string) *sql.DB {
	// Create a new database connection
	db, err := sql.Open(driver, databaseString)
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	return db
}
