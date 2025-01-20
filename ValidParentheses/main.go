package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
    fmt.Println(isValid("()[]{}"))
    fmt.Println(isValid("(]")) 
    fmt.Println(isValid("([])"))
}

func isValid(s string) bool {
	stack := []rune{}
	for _, v := range s {
		if v == '(' {
            stack = append(stack, ')')
        } else if v == '[' {
            stack = append(stack, ']')
        } else if v == '{' {
            stack = append(stack, '}')
        } else {
            if len(stack) == 0 || stack[len(stack)-1] != v {
                return false
            }
            stack = stack[:len(stack)-1]
        }
	}
	return len(stack) == 0
}