package main

import (
	"fmt"
)

func main() {
	array := []int{1, 2}
	fmt.Println(array)
	fmt.Println(topKFrequent(array, 2))
	// fmt.Println(radixSort([]int{4,7,1,9,5}))
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, item := range nums {
		m[item]++
	}
	mm := make(map[int][]int)
	for k, v := range m {
		if slice, ok := mm[v]; !ok {
			mm[v] = []int{k}
		} else {
			mm[v] = append(slice, k)
		}
	}
	result := []int{}
	for k, _ := range mm {
		result = append(result, k)
	}
	fmt.Println("======================= result: ", result)
	final := []int{}
	s := radixSort(result)
	fmt.Println("s: ", s)
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println("s[i] = ", mm[s[i]])
		final = append(final, mm[s[i]]...)
	}
	return final[:k]
}

func radixSort(array []int) []int {
	largestNum := findLargestNum(array)
	size := len(array)
	significantDigit := 1
	semiSorted := make([]int, size, size)
	for largestNum/significantDigit > 0 {
		bucket := [10]int{0}
		for i := 0; i < size; i++ {
			bucket[(array[i]/significantDigit)%10]++
		}
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		for i := size - 1; i >= 0; i-- {
			bucket[(array[i]/significantDigit)%10]--
			semiSorted[bucket[(array[i]/significantDigit)%10]] = array[i]
		}
		for i := 0; i < size; i++ {
			array[i] = semiSorted[i]
		}
		significantDigit *= 10
	}
	return array
}

func findLargestNum(array []int) int {
	largestNum := 0
	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}
