package main

import "fmt"

type Person struct {
	name       string
	age        int
	isLearning bool
}

func main() {
	guy := Person{
		name:       "Allen Robin Hubert",
		age:        29,
		isLearning: false,
	}
	learningStatus := "I'm learning GO."
	if !guy.isLearning {
		learningStatus = "I'm NOT learning GO."
	}
	fmt.Printf("Hi! My name is %s and I'm %d years old. %s", guy.name, guy.age, learningStatus)

}
