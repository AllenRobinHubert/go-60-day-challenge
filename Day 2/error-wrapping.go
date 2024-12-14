package main

import (
	"errors"
	"fmt"
)

var missingConnectionErr = errors.New("missing connection string")

func connectDB(connectionString string) error {
	if connectionString == "" {
		return fmt.Errorf("error %w", missingConnectionErr)
	}
	return nil
}

func initializeServer(connectionString string) error {
	err := connectDB(connectionString)

	if err != nil {
		return fmt.Errorf("server initialization failed: %w", err)
	}

	return nil
}

func main() {
	result := initializeServer("")
	fmt.Println(result)
	if errors.Is(result, missingConnectionErr) {
		fmt.Println("Check: Missing database connection")
	}
}
