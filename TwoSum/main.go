package main

import "fmt"

var arr = []int{1, 2, 3, 4, 5}
var target = 5

func main() {
	sumArr := twoSum()
	fmt.Println(sumArr)
}

func twoSum() []int {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			if arr[i] + arr[j] == target {
				return []int{i, j};
			}
		}
	}
	return nil
}