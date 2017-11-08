package main

import "fmt"

func detectCapitalUse(word string) bool {
	if word == "" {
		return false
	}

	allUpperCaseLeft := true
	allLowerCaseLeft := true
	for i, w := range word {
		if i == 0 {
			continue
		}
		upperCase := 'A' <= w && w <= 'Z'
		lowerCase := 'a' <= w && w <= 'z'
		if !upperCase && !lowerCase {
			return false
		}
		allUpperCaseLeft = allUpperCaseLeft && upperCase
		allLowerCaseLeft = allLowerCaseLeft && lowerCase
	}
	firstUpperCase := 'A' <= word[0] && word[0] <= 'Z'
	firstLowerCase := 'a' <= word[0] && word[0] <= 'z'

	return (firstUpperCase && (allUpperCaseLeft || allLowerCaseLeft)) || (firstLowerCase && allLowerCaseLeft)
}

func main() {
	fmt.Println(detectCapitalUse("USA"))
}
