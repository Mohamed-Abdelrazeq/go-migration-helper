package logs

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-migration-helper/constants"
	"io"
	"log"
	"os"
)

type Stack struct {
	Elements []string `json:"elements"`
}

func readStackFromFile() (*Stack, error) {
	jsonFile, err := os.Open(constants.LogsFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var stack Stack
	if len(data) == 0 {
		return &stack, nil
	}
	err = json.Unmarshal(data, &stack)
	if err != nil {
		return nil, err
	}
	return &stack, nil
}

func writeStackToFile(stack *Stack) error {
	jsonFile, err := os.OpenFile("logs.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	data, err := json.Marshal(stack)
	if err != nil {
		return err
	}

	_, err = jsonFile.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func Push(element string) error {
	stack, err := readStackFromFile()
	if err != nil {
		return errors.New("error reading stack from file")
	}
	stack.Elements = append(stack.Elements, element)
	return writeStackToFile(stack)
}

func Pop() (string, error) {
	stack, err := readStackFromFile()
	if err != nil {
		return "", err
	}
	if len(stack.Elements) == 0 {
		return "", errors.New("logs are empty")
	}
	element := stack.Elements[len(stack.Elements)-1]
	stack.Elements = stack.Elements[:len(stack.Elements)-1]
	err = writeStackToFile(stack)
	if err != nil {
		return "", err
	}
	return element, nil
}
