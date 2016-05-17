package main

import (
	"fmt"
)

func divide(dividend int, divisor int) int {
	divideFunc := func(dividend int, divisor int) int {
		quotient := 0
		if divisor == 1 {
			return dividend
		}
		for dividend >= divisor {
			dividend -= divisor
			quotient++
		}
		return quotient
	}
	if dividend == -2147483648 && divisor == -1 {
		return 2147483647
	}
	if dividend > 0 && divisor > 0 {
		return divideFunc(dividend, divisor)
	} else if dividend < 0 && divisor < 0 {
		return divideFunc(-dividend, -divisor)
	} else if dividend > 0 && divisor < 0 {
		return -divideFunc(dividend, -divisor)
	} else {
		return -divideFunc(-dividend, divisor)
	}
	return 0
}

func main() {
	fmt.Println(divide(5, 2))
	fmt.Println(divide(-5, -2))
	fmt.Println(divide(5, -2))
	fmt.Println(divide(-5, 2))
	fmt.Println(divide(-2147483648, 1))
}
