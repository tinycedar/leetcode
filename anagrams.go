package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}
	if len(strs) > 0 {
		m := make(map[string][]string)
		for _, str := range strs {
			temp := strings.Split(str, "")
			sort.Strings(temp)
			k := strings.Join(temp, "")
			if _, ok := m[k]; !ok {
				m[k] = []string{}
			}
			m[k] = append(m[k], str)
		}
		for _, v := range m {
			sort.Strings(v)
			result = append(result, v)
		}
	}
	return result
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
