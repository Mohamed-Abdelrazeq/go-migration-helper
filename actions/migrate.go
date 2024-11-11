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

func contains(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func Migrate(db *sql.DB) {
	migrationLogs, err := logs.Logs()
	if err != nil {
		log.Fatal("Error reading logs: ", err)
	}

	// Read all files in the migrations folder
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal("Error reading migrations folder: ", err)
	}

	for _, file := range files {
		if !file.IsDir() && !contains(migrationLogs.Elements, file.Name()) {
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
