package actions

import (
	"fmt"
	"go-migration-helper/constants"
	"log"
	"os"
)

func InitMigrationFolder() {
	log.Println("Starting the application...")

	// Creating logs file
	if _, err := os.Stat("logs.json"); os.IsNotExist(err) {
		file, err := os.Create("logs.json")
		if err != nil {
			log.Fatal("Error creating logs.json file: ", err)
		}
		defer file.Close()
	}

	// Create a migrations folder
	err := os.Mkdir("migrations", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}

	// Create a initial migration file
	file, err := os.Create("migrations/0001_initial.sql")
	if err != nil {
		log.Fatal("Error creating migration file: ", err)
	}
	defer file.Close()

	_, err = file.WriteString(constants.Template)
	if err != nil {
		log.Fatal("Error writing to migration file: ", err)
	}

	fmt.Println("Project initialized successfully!")
}
