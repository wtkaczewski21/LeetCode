package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s))
}

func lengthOfLastWord(s string) int {
    split := strings.Fields(s)
	return len(split[len(split)-1])
}