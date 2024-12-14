package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create(("output.text"))
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure the file is closed

	for i := 1; i <= 10000000; i++ {
		// Write data to the file
		message := "line - " + strconv.Itoa(i) + "\n"
		_, err = file.WriteString(message)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("File created and written successfully!")
}
