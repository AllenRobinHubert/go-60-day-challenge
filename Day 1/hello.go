package main

import "fmt"

func main() {

	var name string = "Allen Robin Hubert"
	var age int = 29
	var isLearning bool = false

	learningStatus := "I'm learning GO."

	if !isLearning {
		learningStatus = "I'm NOT learning GO."
	}

	fmt.Printf("Hi! My name is %s and I'm %d years old. %s", name, age, learningStatus)

}
