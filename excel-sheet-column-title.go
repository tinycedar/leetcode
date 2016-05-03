package main

import (
	"fmt"
)

func convertToTitle(n int) string {
	m := make(map[int]rune)
	for i := 0; i < 26; i++ {
		m[i+1] = rune('A' + i)
	}
	slice := []rune{}
	if n <= 26 {
		return string(m[n])
	}
	for ; n > 0; n /= 26 {
		if n%26 == 0 {
			slice = append(slice, m[26])
			n -= 26
			for n > 0 && n <= 26 {
				n /= 26
				if n > 0 {
					slice = append(slice, m[n])
				}
			}
		} else {
			slice = append(slice, m[n%26])
		}
	}

	sliceLen := len(slice)
	for i, j := 0, sliceLen-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}

func main() {
	// fmt.Println("26= ", convertToTitle(26))
	// fmt.Println("27= ", convertToTitle(27))
	// fmt.Println("28= ", convertToTitle(28))
	// fmt.Println("52= ", convertToTitle(52))
	// fmt.Println("701= ", convertToTitle(701))
	fmt.Println("702= ", convertToTitle(702))
}
