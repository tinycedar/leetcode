package main

import (
	"fmt"
	"bytes"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) String() string {
	var buf = bytes.Buffer{}
	for n := head; n != nil; n = n.Next {
		if n != head {
			buf.WriteString(" -> ")
		}
		buf.WriteString(fmt.Sprintf("[%d]", n.Val))
	}

	return buf.String()
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, newTail, oldHead, oldTail *ListNode
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val >= x {
			if oldHead == nil {
				oldHead = cur
			}
			oldTail = cur
		} else {
			if newHead == nil {
				newHead, newTail = cur, cur
			} else {
				newTail.Next = cur
				newTail = cur
			}
			if oldTail != nil {
				oldTail.Next = cur.Next
			}
		}
	}
	if newHead == nil {
		return head
	}
	newTail.Next = oldHead
	return newHead
}

func main() {
	h := &ListNode{1, &ListNode{1, nil}}
	fmt.Println(h)
	fmt.Println(partition(h, 2))
}
