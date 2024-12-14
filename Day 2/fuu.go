package main

import (
	"fmt"
)

type calculationError struct {
	Operation    string
	FirstNumber  int
	SecondNumber int
	Message      string
}

func (e *calculationError) Error() string {
	return fmt.Sprintf("Error during %s operation : %s. Numbers %d, %d. ", e.Operation, e.Message, e.FirstNumber, e.SecondNumber)
}

func calculate(a int, b int, operation string) (int, error) {

	var result int

	switch operation {
	case "divide":
		if a == 0 || b == 0 {
			return 0, &calculationError{
				Operation:    operation,
				FirstNumber:  a,
				SecondNumber: b,
				Message:      "Cannot devide with 0",
			}
		}
		result = a / b

	case "add":
		result = a + b

	case "subtract":
		result = a - b

	case "multiply":
		result = a * b

	default:
		return 0, &calculationError{
			Operation:    operation,
			FirstNumber:  a,
			SecondNumber: b,
			Message:      "Invalid operation",
		}
	}

	return result, nil
}

func main() {
	result, err := calculate(100, 20, "divide")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
