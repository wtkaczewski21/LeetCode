package main

import "fmt"

func main() {
	digits := []int{1, 9, 9}
	fmt.Println(plusOne(digits))
}

func plusOne(digits []int) []int {
	for last := len(digits) - 1; last >= 0; last-- {
		if digits[last] != 9 {
			digits[last] += 1
			break
		}
		digits[last] = 0
	}
	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}