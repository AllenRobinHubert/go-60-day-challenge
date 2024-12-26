package main

import "fmt"

func main() {
	var randomArr = [6]int{1, 2, 3, 4, 5, 6}
	reOrderedSlice := []int{}

	for _, v := range randomArr {
		reOrderedSlice = append(reOrderedSlice, v)
	}

	reOrderedSlice = append(reOrderedSlice, 2, 3, 4)

	reOrderedSlice = append(reOrderedSlice[:1], reOrderedSlice[2:]...)

	fmt.Println(reOrderedSlice)

}
