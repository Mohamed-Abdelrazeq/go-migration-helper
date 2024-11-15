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

func Rollback(db *sql.DB) {
	latestMigration, err := logs.Pop()
	if err != nil {
		log.Fatal("No migration to rollback")
	}

	// Read latest migration file from the migrations folder
	filePath := "migrations/" + latestMigration
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
		log.Fatal("Error extracting down migration: ", err)
	}
	_, err = db.Exec(downMigration)
	if err != nil {
		log.Fatal("Error executing migration: ", err)
	}

	fmt.Printf(`Miration %s rolled back successfully!`, latestMigration)
}
