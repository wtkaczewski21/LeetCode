package main

import "fmt"

func main() {
	var s = "III"
	fmt.Println(romanToInt(s))
}

func value(s byte) int {
	if s == 'I' {
		return 1
	}
	if s == 'V' {
		return 5
	}
	if s == 'X' {
		return 10
	}
	if s == 'L' {
		return 50
	}
	if s == 'C' {
		return 100
	}
	if s == 'D' {
		return 500
	}
	if s == 'M' {
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		s1 := value(s[i])
		if i + 1 < len(s) {
			s2 := value(s[i +1])
			if s1 >= s2 {
				res += s1
			} else {
				res += (s2 - s1)
				i++
			}
		} else {
			res += s1
		}
	}
	return res
}