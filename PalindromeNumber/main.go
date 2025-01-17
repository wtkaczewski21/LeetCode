package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	number := 121
	if isPalindrome(number) {
		fmt.Println("Number is palindrome")
	} else {
		fmt.Println("Number is not palindrome")
	}
}
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	strArr := strings.Split(str, "")
	var reversed []string
    for i := len(strArr)-1; i >= 0; i-- {
		reversed = append(reversed, strArr[i])
    }
	return reflect.DeepEqual(reversed, strArr)
}