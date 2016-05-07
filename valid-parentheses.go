package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	for _, item := range s {
		size := len(stack)
		if size == 0 {
			stack = append(stack, item)
		} else {
			// top
			top := stack[size-1]
			if top == '(' && item == ')' || top == '{' && item == '}' || top == '[' && item == ']' {
				// pop
				stack = stack[:size-1]
			} else {
				// push
				stack = append(stack, item)
			}
		}
		// fmt.Println("stack= ", stack)
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	// fmt.Println(isValid("([{}])"))
}
