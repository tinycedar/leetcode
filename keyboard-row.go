package main

import (
	"strings"
	"bytes"
)

func findWords(words []string) []string {
	if words == nil || len(words) == 0 {
		return []string{}
	}
	firstLineMap := map[rune]bool{
		'q': true, 'w': true, 'e': true, 'r': true, 't': true, 'y': true, 'u': true, 'i': true, 'o': true, 'p': true,
		'Q': true, 'W': true, 'E': true, 'R': true, 'T': true, 'Y': true, 'U': true, 'I': true, 'O': true, 'P': true,
	}
	secondLineMap := map[rune]bool{
		'a': true, 's': true, 'd': true, 'f': true, 'g': true, 'h': true, 'j': true, 'k': true, 'l': true,
		'A': true, 'S': true, 'D': true, 'F': true, 'G': true, 'H': true, 'J': true, 'K': true, 'L': true,
	}

	thirdLineMap := map[rune]bool{
		'z': true, 'x': true, 'c': true, 'v': true, 'b': true, 'n': true, 'm': true,
		'Z': true, 'X': true, 'C': true, 'V': true, 'B': true, 'N': true, 'M': true,
	}
	var buf = bytes.Buffer{}
	for _, w := range words {
		if containsAll(w, firstLineMap, secondLineMap, thirdLineMap) {
			if buf.Len() > 0 {
				buf.WriteString("_")
			}
			buf.WriteString(w)
		}
	}
	if buf.Len() > 0 {
		return strings.Split(buf.String(), "_")
	}
	return []string{}
}

func containsAll(word string, firstLineMap, secondLineMap, thirdLineMap map[rune]bool) bool {
	if word == "" {
		return false
	}
	first := true
	second := true
	third := true
	for _, w := range word {
		if first {
			if ok, _ := firstLineMap[w]; !ok {
				first = false
			}
		}
		if second {
			if ok, _ := secondLineMap[w]; !ok {
				second = false
			}
		}
		if third {
			if ok, _ := thirdLineMap[w]; !ok {
				third = false
			}
		}
	}
	return first || second || third
}

func main() {
	print(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}

func print(words []string) {
	for _, w := range words {
		println(w)
	}
}
