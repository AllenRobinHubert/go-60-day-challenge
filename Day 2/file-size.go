package main

import (
	"fmt"
)

func uploadFile(fileName string, fileSize int) (string, error) {

	var sizeError string
	var nameError string

	if fileName == "" {
		nameError = fmt.Sprintf("Name Error. File name cannot be empty")
	}
	if fileSize > 100 {
		sizeError = fmt.Sprintf("Size Error. File size cannot be more than 100")
	}

	if sizeError == "" && nameError == "" {
		return "success", nil
	}

	return "failed", fmt.Errorf("Error %s %s", nameError, sizeError)
}

func processUpload(fileName string, fileSize int) (string, error) {
	status, error := uploadFile(fileName, fileSize)

	if error == nil {
		return status, error
	} else {
		return status, fmt.Errorf("Error uploading file: %w", error)
	}
}

func main() {

	fmt.Println("\n")
	status, error := processUpload("test", 200)

	if error == nil {
		fmt.Printf("Status: %s ", status)
	} else {
		fmt.Printf("Status: %s \n Error: %w", status, error)
	}

	fmt.Println("\n")
}
