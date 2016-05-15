package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	str := ""
	for cur := head; cur != nil; cur = cur.Next {
		str += fmt.Sprintf("->%d", cur.Val)
	}
	return str
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	m := make(map[int]int)
	for cur := head; cur != nil; cur = cur.Next {
		m[cur.Val]++
	}
	for ; head != nil && m[head.Val] > 1; head = head.Next {
		m[head.Val]--
	}
	if head == nil || head.Next == nil {
		return head
	}
	for pre, cur := head, head.Next; cur != nil; {
		if m[cur.Val] > 1 {
			m[cur.Val]--
			pre.Next, cur = cur.Next, cur.Next
		} else {
			pre, cur = pre.Next, cur.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{2, nil}}}
	fmt.Println(head)
	fmt.Println(deleteDuplicates(head))
}
