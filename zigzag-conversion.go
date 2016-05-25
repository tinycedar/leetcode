package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows <= 1 {
		return s
	}
	matrix := [][]byte{}
	var inner []byte
	for i := 0; i < len(s); {
		if len(inner) <= 0 {
			inner = make([]byte, 0)
		}
		isZig := numRows > 2 && len(matrix) > 0 && len(matrix)%2 != 0
		if len(inner) == numRows || isZig && len(inner) == numRows-2 {
			matrix = append(matrix, inner)
			inner = nil
			continue
		}
		inner = append(inner, s[i])
		if i == len(s)-1 {
			if isZig {
				zeroToAdd := numRows - 2 - len(inner)
				for j := 0; j < zeroToAdd; j++ {
					inner = append(inner, byte(0))
				}
			}
			matrix = append(matrix, inner)
		}
		i++
	}
	for i, v := range matrix {
		if isZig := i%2 != 0 && numRows > 2; isZig {
			if len(v) > 1 {
				reverse(v)
			}
			temp := []byte{}
			temp = append(temp, 0)
			temp = append(temp, v...)
			temp = append(temp, 0)
			v = temp
		}
		if len(v) < numRows {
			zeroToAdd := numRows - len(v)
			for j := 0; j < zeroToAdd; j++ {
				v = append(v, 0)
			}
		}
		matrix[i] = v
	}
	result := make([]string, len(s))
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(matrix); j++ {
			if item := matrix[j][i]; item != 0 {
				result = append(result, string(item))
			}
		}
	}
	return strings.Join(result, "")
}

func reverse(slice []byte) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func main() {
	// fmt.Println(convert("ABC", 2))
	// fmt.Println(convert("ABCD", 2))
	fmt.Println(convert("ABCDE", 4))
}
