package actions

import (
	"database/sql"
	"fmt"
	"go-migration-helper/helpers"
	"io"
	"log"
	"os"
)

func ResetMigrations(db *sql.DB) {
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
			downMigration, err := helpers.ExtractUpOrDownMigration("down", string(content))
			if err != nil {
				log.Fatal("Error extracting up migration: ", err)
			}
			_, err = db.Exec(downMigration)
			if err != nil {
				log.Fatal("Error executing migration: ", err)
			}

			fmt.Printf(`Miration %s executed successfully!`, file.Name())
			fmt.Println()
		}
	}
}
