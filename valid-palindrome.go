package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	length := len(s)
	if length <= 1 {
		return true
	}
	for i, j := 0, length-1; i <= j; {
		a := s[i]
		b := s[j]
		if !isAlphanumeric(a) {
			i++
			continue
		}
		if !isAlphanumeric(b) {
			j--
			continue
		}
		if isAlphanumeric(a) && isAlphanumeric(b) {
			if !isEqualIgnoreCase(a, b) {
				return false
			}
		}
		i, j = i+1, j-1
	}
	return true
}

func isEqualIgnoreCase(a, b byte) bool {
	return strings.ToLower(string(a)) == strings.ToLower(string(b))
}

func isAlphanumeric(c byte) bool {
	v := strings.ToLower(string(c))
	if v >= "a" && v <= "z" ||
		v >= "0" && v <= "9" {
		return true
	}
	return false
}

func main() {
	// fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("!bHvX!?!!vHbX"))
}
