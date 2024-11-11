package actions

import (
	"database/sql"
	"fmt"
	"go-migration-helper/helpers"
	"go-migration-helper/logs"
	"io"
	"log"
	"os"
)

func Migrate(db *sql.DB) {
	// Read all files in the migrations folder
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal("Error reading migrations folder: ", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := "migrations/" + file.Name()
			f, err := os.Open(filePath)
			if err != nil {
				log.Fatal("Error opening file: ", err)
			}

			content, err := io.ReadAll(f)
			if err != nil {
				log.Fatal("Error reading file: ", err)
			}
			defer f.Close()

			// Execute the migration
			upMigration, err := helpers.ExtractUpOrDownMigration("up", string(content))
			if err != nil {
				log.Fatal("Error extracting up migration: ", err)
			}
			_, err = db.Exec(upMigration)
			if err != nil {
				log.Fatal("Error executing migration: ", err)
				return
			}
			logs.Push(file.Name())

			fmt.Printf("Migration %s executed successfully!\n", file.Name())
		}
	}
}
