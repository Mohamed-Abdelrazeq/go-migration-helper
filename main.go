package main

// postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable
import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const newMigrationTemplate = `-- +migrate Up

-- -migrate Down
		`

func main() {
	log.Println("Starting the application...")

	if len(os.Args) < 2 {
		log.Fatal("Expected 'up' or 'down' command")
	}

	command := os.Args[1]

	driver, databaseString := scanDatabaseInfo()

	db := initConnection(driver, databaseString)
	defer db.Close()

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
	case "add":
		addMigration()
	case "migrate":
		migrateDatabase(db)
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

	// Create .env file if it doesn't exist
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		file, err := os.Create(".env")
		if err != nil {
			log.Fatal("Error creating .env file: ", err)
		}
		defer file.Close()
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

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
	envFile, err := os.OpenFile(".env", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal("Error opening .env file: ", err)
	}
	defer envFile.Close()

	_, err = envFile.WriteString(fmt.Sprintf("GMH_DB_DRIVER=%s\nGMH_DB_STRING=%s\n", driver, databaseString))
	if err != nil {
		log.Fatal("Error writing to .env file: ", err)
	}

	return driver, databaseString
}

func initConnection(driver string, databaseString string) *sql.DB {
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
	file, err := os.Create("migrations/0001_initial.sql")
	if err != nil {
		log.Fatal("Error creating migration file: ", err)
	}
	defer file.Close()

	_, err = file.WriteString(newMigrationTemplate)
	if err != nil {
		log.Fatal("Error writing to migration file: ", err)
	}

	fmt.Println("Migration file created successfully!")
}

func migrateDatabase(db *sql.DB) {
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
			_, err = db.Exec(string(content))
			if err != nil {
				log.Fatal("Error executing migration: ", err)
			}

			fmt.Printf(`Miration %s executed successfully!`, file.Name())
			fmt.Println()
		}
	}
}

func addMigration() {

	// Get the new migration file name from args
	if len(os.Args) < 3 {
		log.Fatal("Expected migration file name")
	}

	newFileName := os.Args[2]

	// Scan most recent migration file
	files, err := os.ReadDir("migrations")
	if err != nil {
		log.Fatal("Error reading migrations folder: ", err)
	}

	lastFileName := files[len(files)-1].Name()
	lastFileNumber, err := strconv.Atoi(lastFileName[:4])
	if err != nil {
		log.Fatal("Error converting file number to integer: ", err)
	}

	fmt.Println(lastFileNumber)

	newFileFullName := fmt.Sprintf("migrations/%s_%s.sql", fmt.Sprintf("%04d", lastFileNumber+1), newFileName)

	println(newFileFullName)
	// Create a new migration file
	file, err := os.Create(newFileFullName)
	if err != nil {
		log.Fatal("Error creating migration file: ", err)
	}
	defer file.Close()

	_, err = file.WriteString(newMigrationTemplate)
	if err != nil {
		log.Fatal("Error writing to migration file: ", err)
	}

	fmt.Println("Migration file created successfully!")
}
