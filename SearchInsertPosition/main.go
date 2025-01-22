package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	// target := 2
	// target := 7
	fmt.Println(searchInsert(nums, target))
}

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		}
		if v > target {
			return i
		}
	}
	return len(nums)
}