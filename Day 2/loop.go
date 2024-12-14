package main

import "fmt"

func main() {

	for i := 1; i <= 20; i++ {

		if i == 13 || i == 17 {
			message := fmt.Sprintf("Skipping (%d)", i)
			fmt.Println(message)
			continue
		}

		if i%5 == 0 {
			message := fmt.Sprintf("This number (%d) is divisible by 5!", i)
			fmt.Println(message)
		}

		if i%2 == 0 {
			message := fmt.Sprintf("The number (%d) is even!", i)
			fmt.Println(message)
		} else {
			message := fmt.Sprintf("The number (%d) is odd!", i)
			fmt.Println(message)
		}
	}

}
