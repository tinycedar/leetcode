package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	list := ""
	for temp := head; temp != nil; temp = temp.Next {
		if temp == head {
			list = fmt.Sprintf("%d", temp.Val)
		} else {
			list = fmt.Sprintf("%s->%d", list, temp.Val)
		}
	}
	return list
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head != nil && head.Next != nil {
		next := head.Next
		head.Next = next.Next
		next.Next = head
		head = next
	}
	for previous, current := head.Next, head.Next.Next; previous != nil && current != nil && current.Next != nil; {
		next := current.Next
		previous.Next = next
		if next != nil {
			current.Next = next.Next
			next.Next = current
		}
		previous = current
		current = current.Next
	}
	return head
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(swapPairs(list))
	list = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	fmt.Println(swapPairs(list))
}
