package main

import (
	"errors"
	"fmt"
)

var errZeroDivisible = errors.New("cannot be divided by zero")

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("error %w", errZeroDivisible)
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
