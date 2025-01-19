package main

import (
	"fmt"
	"sort"
)

func main() {
	// strs := []string{"dog", "racecar", "car"}
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
    	return strs[i] < strs[j]
	})
	first := strs[0]
	last := strs[len(strs)-1]
	minLength := min(first, last)
	var bytes []byte
	for i := 0; i < len(minLength); i++ {
		if first[i] != last[i] {
			break
		}
		bytes = append(bytes, first[i])
	}
	prefix := string(bytes)
    return prefix
}