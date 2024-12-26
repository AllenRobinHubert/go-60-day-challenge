package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func isEven(x int) bool {
	var isTrue bool
	if x%2 == 0 {
		isTrue = true
	} else {
		isTrue = false
	}
	return isTrue
}

func main() {
	sumResult := sum(30, 40)
	fmt.Println(sumResult)
	fmt.Println(isEven(29))
}
