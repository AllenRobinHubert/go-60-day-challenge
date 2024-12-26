package main

import (
	"fmt"

	"github.com/fatih/color"
)

func names() {
	names := [5]string{"allen", "robin", "hubert", "nibi", "mathews"}
	names[1] = "john (Updated)"
	fmt.Println("Updated Student Names:")
	for i, name := range names {
		message := fmt.Sprintf("%d - %s", i+1, name)
		fmt.Println(message)
	}
}

func scores() {
	scores := []int{85, 90, 75}
	scores = append(scores, 13, 28)

	scores = append(scores[:3], scores[4:]...)
	fmt.Println("Updated Student Scores:")
	for _, score := range scores {
		fmt.Println(score)
	}
}

func namesandscores() {
	scores := make(map[string]int)

	scores["john"] = 95
	scores["charlie"] = 75
	scores["david"] = 88
	scores["john"] = 99
	scores["Alice"] = 59
	scores["sherin"] = 40
	scores["bob"] = 40

	value, exists := scores["Alice"]
	if exists {
		fmt.Println("Alice's Score:", value)
	} else {
		fmt.Println("Alice not found")
	}

	fmt.Println("")

	fmt.Println("Student Scores:")
	delete(scores, "Bob")

	for k, v := range scores {
		message := fmt.Sprintf("%s - %d", k, v)
		fmt.Println(message)
	}

}

func main() {

	fmt.Println("")
	names()
	fmt.Println("")
	scores()
	fmt.Println("")
	namesandscores()
	fmt.Println("")
	color.Cyan("Hello, Colorful World!")
}
