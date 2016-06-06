//Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].

package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) <= 0 {
		return []int{}
	}
	m := make(map[int]int)
	for _, v := range nums2 {
		m[v]++
	}
	resultMap := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			resultMap[v]++
		}
	}
	result := []int{}
	for k := range resultMap {
		result = append(result, k)
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))

	// abc := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }

	// keys := reflect.ValueOf(abc).MapKeys()

	// fmt.Println(keys) // [a b c]
}
