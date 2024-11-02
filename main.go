package main

// postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
import (
	"database/sql"
	"fmt"
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

	driver, databaseString := scanDatabaseInfo()

	_ = initConnection(driver, databaseString)

	// TODO: Read all files in the migrations folder
	// fileName := "migrations/001_initial.sql"

	// Read the migration file
	// content, err := io.ReadAll(fileName)
	// if err != nil {
	// 	log.Fatal("Error reading migration file: ", err)
	// }

	switch command {
	case "init":
		initMigrationFolder()
	case "migrate":
		log.Fatal("Migrate not implemented yet")
	case "rollback":
		log.Fatal("Rollback not implemented yet")
	case "reset":
		log.Fatal("Reset not implemented yet")
	default:
		log.Fatal("Unknown command: ", command)
	}

}

func scanDatabaseInfo() (string, string) {
	var driver string
	var databaseString string

	// Check if the database info is cached
	if os.Getenv("GMH_DB_DRIVER") != "" && os.Getenv("GMH_DB_STRING") != "" {
		return os.Getenv("GMH_DB_DRIVER"), os.Getenv("GMH_DB_STRING")
	}

	// Scan the database info
	fmt.Println("Enter database driver:")
	fmt.Scanln(&driver)
	fmt.Println("Enter database string:")
	fmt.Scanln(&databaseString)

	// Cache the database info
	os.Setenv("GMH_DB_DRIVER", driver)
	os.Setenv("GMH_DB_STRING", databaseString)

	return driver, databaseString
}

func initConnection(driver string, databaseString string) *sql.DB {
	// Create a new database connection
	db, err := sql.Open(driver, databaseString)
	if err != nil {
		log.Fatal("Error opening the database: ", err)
	}
	defer db.Close()

	// Verify the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}

func initMigrationFolder() {
	err := os.Mkdir("migrations", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	fmt.Println("Migration folder created successfully!")

	// Create a initial migration file
	file, err := os.Create("migrations/001_initial.sql")
	if err != nil {
		log.Fatal("Error creating migration file: ", err)
	}
	defer file.Close()

	fmt.Println("Migration file created successfully!")
}
