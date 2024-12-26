package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {

	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		var complement = target - nums[i]
		val, ok := hash[complement]
		if ok {
			return []int{val, i}
		}
		hash[nums[i]] = i
	}
	return nil

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	og := x
	reversed := 0

	for x > 0 {

		new_int := x % 10
		reversed = reversed*10 + new_int
		x = x / 10
	}

	if reversed != og {
		return false
	}

	return true
}

func isPalindromeTwo(x int) bool {

	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	reversed := 0

	for x > reversed {
		lastDigit := x % 10
		reversed = reversed*10 + lastDigit
		x = x / 10
	}

	return x == reversed || x == reversed/10
}

func romanToInt(s string) int {

	roman := make(map[string]int)

	roman["I"] = 1
	roman["V"] = 5
	roman["X"] = 10
	roman["L"] = 50
	roman["C"] = 100
	roman["D"] = 500
	roman["M"] = 1000

	var prev string
	var final int = 0
	for _, c := range s {
		// fmt.Printf("%d %c\n", i, c)

		if prev != "" {
			if roman[prev] >= roman[string(c)] {
				// fmt.Println(string(c))
				final = final + roman[string(c)]
			} else {
				fmt.Println("previous is smaller")
			}
		}

		prev = string(c)
	}

	fmt.Println(final)
	return 1

}

func main() {
	// result := twoSum([]int{2, 7, 11, 15}, 9)
	romanToInt("IIII")
	// fmt.Println(result)
}
