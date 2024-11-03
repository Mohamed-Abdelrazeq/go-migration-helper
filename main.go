package main

// postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
import (
	"go-migration-helper/actions"
	"go-migration-helper/helpers"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	log.Println("Starting the application...")

	if len(os.Args) < 2 {
		log.Fatal("Expected 'up' or 'down' command")
	}

	command := os.Args[1]

	// Create .env file if it doesn't exist
	helpers.InitCache()

	db := helpers.Connect(helpers.ScanDatabaseInfo())
	defer db.Close()

	switch command {
	case "init":
		actions.InitMigrationFolder()
	case "add":
		actions.AddMigration()
	case "migrate":
		actions.Migrate(db)
	case "rollback":
		actions.Rollback(db)
	case "reset":
		// TODO: Execute the down migration for all migration files
		actions.ResetMigrations(db)
	default:
		log.Fatal("Unknown command: ", command)
	}

}
