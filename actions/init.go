package actions

import (
	"fmt"
	"go-migration-helper/constants"
	"log"
	"os"
)

func InitMigrationFolder() {
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
