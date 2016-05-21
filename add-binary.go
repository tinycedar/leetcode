package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}
	temp := []string{}
	for _, v := range a[:len(a)-len(b)] {
		temp = append(temp, string(v))
	}
	for i, _ := range b {
		temp = append(temp, string(a[len(a)-len(b):][i]+b[i]-'0'))
	}
	for i := len(temp) - 1; i > 0; i-- {
		cur, _ := strconv.Atoi(temp[i])
		if cur > 1 {
			pre, _ := strconv.Atoi(temp[i-1])
			temp[i-1] = fmt.Sprintf("%d", pre+cur/2)
			temp[i] = fmt.Sprintf("%d", cur%2)
		}
	}
	first, _ := strconv.Atoi(temp[0])
	if first > 1 {
		temp[0] = fmt.Sprintf("%d", first%2)
		newTemp := []string{}
		newTemp = append(newTemp, fmt.Sprintf("%d", first/2))
		newTemp = append(newTemp, temp...)
		temp = newTemp
	}
	return strings.Join(temp, "")
}

func main() {
	// fmt.Println(addBinary("111", "11111"))
	// fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("100", "110010"))
}
