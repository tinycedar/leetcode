package main

import (
	"fmt"
)

func isPowerOfFour(num int) bool {
	if num <= 0 {
		return false
	}
	for ; num%4 == 0; num /= 4 {
	}
	return num == 1
}

func main() {
	fmt.Println(isPowerOfFour(8))
}
