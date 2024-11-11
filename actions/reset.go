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

func ResetMigrations(db *sql.DB) {

	// read logs
	for {
		migrationFileName, err := logs.Pop()
		if err != nil {
			return
		}

		// read migration file
		filePath := "migrations/" + migrationFileName
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
		}

		content, err := io.ReadAll(f)
		if err != nil {
			log.Fatal("Error reading file: ", err)
		}
		f.Close()

		downMigration, err := helpers.ExtractUpOrDownMigration("down", string(content))
		if err != nil {
			log.Fatal("Error extracting up migration: ", err)
		}
		_, err = db.Exec(downMigration)
		if err != nil {
			log.Fatal("Error executing migration: ", err)
		}

		fmt.Printf(`Rolling back %s successfully!`, f.Name())
		fmt.Println()
	}

}
