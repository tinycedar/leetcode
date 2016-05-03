package main

import (
	"fmt"
)

func main() {
	fmt.Println(titleToNumber("ABC"))
}

func titleToNumber(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	m := make(map[rune]int)
	for char := 'A'; char <= 'Z'; char++ {
		m[char] = int(char - 'A' + 1)
	}
	result := 0
	for index, i := 0, length-1; i >= 0; i-- {
		temp := 1
		for j := 0; j < index; j++ {
			temp *= 26
		}
		result += m[rune(s[i])] * temp
		index++
	}
	return result
}
