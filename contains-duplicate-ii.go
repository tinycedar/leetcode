package main

import (
	"fmt"
)

// Given an array of integers and an integer k,
// find out whether there are two distinct indices i and j in the array
// such that nums[i] = nums[j] and the difference between i and j is at most k.
func containsNearbyDuplicate(nums []int, k int) bool {
	length := len(nums)
	if length <= 1 || k <= 0 {
		return false
	}
	valueMap := make(map[int][]int) // v: []i
	for i, v := range nums {
		s, ok := valueMap[v]
		if !ok {
			s = []int{i}
		} else {
			s = append(s, i)
		}
		valueMap[v] = s
	}
	fmt.Println("nums= ", nums)
	for _, v := range valueMap {
		length := len(v)
		if length <= 1 {
			continue
		}
		fmt.Println("==================== v = ", v)
		for i := 0; i < length-1; i++ {
			for j := i + 1; j < length; j++ {
				if findDifference(nums, v[i], v[j]) <= k {
					return true
				}
			}
		}
	}
	return false
}

func findDifference(nums []int, begin int, end int) int {
	result := make(map[int]int)
	for i := begin; i <= end; i++ {
		result[nums[i]]++
	}
	fmt.Println(begin, end, result)
	return len(result)
}

func main() {
	// fmt.Println(containsNearbyDuplicate([]int{99, 99}, 2))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 1, 1}, 1))
}
