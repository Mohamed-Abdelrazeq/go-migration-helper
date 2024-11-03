package actions

import (
	"fmt"
	"go-migration-helper/constants"
	"log"
	"os"
	"strconv"
)

func AddMigration() {

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

	_, err = file.WriteString(constants.Template)
	if err != nil {
		log.Fatal("Error writing to migration file: ", err)
	}

	fmt.Println("Migration file created successfully!")
}
