package main

import (
	"fmt"
)

func isHappy(n int) bool {
	if n == 1 {
		return true
	}
	m := make(map[int]bool)
	for ; n != 1; m[n] = true {
		sum := 0
		for ; n > 0; n /= 10 {
			sum += (n % 10) * (n % 10)
		}
		n = sum
		if _, ok := m[n]; ok {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isHappy(7))
}
