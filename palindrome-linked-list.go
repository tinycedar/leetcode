package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	list := []int{}
	for n := head; n != nil; n = n.Next {
		list = append(list, n.Val)
	}
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		if list[i] != list[j] {
			return false
		}
	}
	return true
}

func main() {
	list := &ListNode{0, &ListNode{0, nil}}
	fmt.Println(isPalindrome(list))
}
