package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("abc   "))
	fmt.Println(lengthOfLastWord("   abc"))
}

func lengthOfLastWord(s string) int {
	length := len(s)
	spacePos := -1
	nonSpacePos := -1
	for i := length - 1; i >= 0; i-- {
		if s[i] != ' ' {
			if nonSpacePos < 0 {
				nonSpacePos = i
			}
		} else if nonSpacePos >= 0 {
			spacePos = i
			break
		}
	}
	return nonSpacePos - spacePos
}
