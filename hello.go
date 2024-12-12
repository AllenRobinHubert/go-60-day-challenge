package main

import "fmt"

func main() {

	var name string = "Allen Robin Hubert"
	var age int = 29
	var isLearning bool = false

	learningStatus := "im learning go"
	if !isLearning {
		learningStatus = "im NOT learning go"
	}

	message := fmt.Sprintf("My name is %s, im %d years old and %s", name, age, map[bool]string{true: "im learning go", false: "imNOTlearninggo"}[isLearning])
	fmt.Println(message)
	fmt.Printf("My name is %s, im %d years old and %s", name, age, learningStatus)

}
