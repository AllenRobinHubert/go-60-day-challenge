package main

import "fmt"

func main() {
	employees := make(map[string]int)
	employees["allen"] = 78000
	employees["sherin"] = 23000
	employees["nibi"] = 45000
	employees["aneesh"] = 23000
	employees["ashwin"] = 23000

	employees["aneesh"] = 43000

	delete(employees, "nibi")

	for key, value := range employees {
		fmt.Println(key, value)
	}
}
