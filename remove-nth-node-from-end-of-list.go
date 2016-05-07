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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil && n == 1 {
		return nil
	}
	if n <= 0 {
		return head
	}
	count := 0
	for temp := head; temp != nil; temp = temp.Next {
		count++
	}
	if count >= n {
		index := count - n + 1
		if index == 1 {
			return head.Next
		}
		currentPos := 2
		for previous, current := head, head.Next; current != nil; previous, current = previous.Next, current.Next {
			if currentPos == index {
				previous.Next = current.Next
				break
			}
			currentPos++
		}
	}
	return head
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	// fmt.Println(removeNthFromEnd(list, 1))
	// fmt.Println(removeNthFromEnd(list, 2))
	// fmt.Println(removeNthFromEnd(list, 3))
	fmt.Println(removeNthFromEnd(list, 4))
}
