package main

import "fmt"

func main() {
	str1 := "1010"
	str2 := "1011"
	fmt.Println(addBinary(str1, str2))
}

func addBinary(a string, b string) string {
	if len(a) > len(b) {
		return addBinary(b, a)
	}

	difference := len(b) - len(a)

	for i := 0; i < difference; i++ {
	a = "0" + a
	}

	carry := "0"
	answer := ""

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '1' && b[i] == '1' {
			if carry == "1" {
				answer = "1" + answer
			} else {
				answer = "0" + answer
				carry = "1"
			}
		} else if a[i] == '0' && b[i] == '0' {
			if carry == "1" {
				answer = "1" + answer
				carry = "0"
			} else {
				answer = "0" + answer
			}
		} else if a[i] != b[i] {
			if carry == "1" {
				answer = "0" + answer
			} else {
				answer = "1" + answer
			}
		}
	}
	if carry == "1" {
		answer = "1" + answer
	}
	return answer
}