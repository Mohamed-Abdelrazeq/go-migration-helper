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
	migrationLogs, err := logs.Logs()
	if err != nil {
		log.Fatal("Error reading logs: ", err)
	}
	for i := 0; i < len(migrationLogs.Elements); i++ {
		migrationFileName, err := logs.Pop()
		if err != nil {
			log.Fatal("Error reading logs: ", err)
			break
		}

		// read migration file
		filePath := "migrations/" + migrationFileName
		f, err := os.Open(filePath)
		if err != nil {
			log.Fatal("Error opening file: ", err)
			break
		}

		content, err := io.ReadAll(f)
		if err != nil {
			log.Fatal("Error reading file: ", err)
			break
		}
		f.Close()

		downMigration, err := helpers.ExtractUpOrDownMigration("down", string(content))
		if err != nil {
			log.Fatal("Error extracting up migration: ", err)
			break
		}
		_, err = db.Exec(downMigration)
		if err != nil {
			log.Fatal("Error executing migration: ", err)
			break
		}

		fmt.Printf(`Rolled back %s excuted successfully!`, f.Name())
		fmt.Println()
	}

}
