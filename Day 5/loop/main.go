package main

import "fmt"

func main() {
	var limit int = 50
	for i := 0; i < limit; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Buzz")
		} else if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}

}
