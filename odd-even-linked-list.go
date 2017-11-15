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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 0
	for n := head; n != nil; n = n.Next {
		length ++
	}
	var evenHead, evenTail *ListNode
	lastOdd, node := head, head
	for i := 1; i <= length; i++ {
		if i%2 == 0 {
			if evenHead == nil {
				evenHead, evenTail = node, node
			} else {
				evenTail.Next = node
				evenTail = node
			}
			lastOdd.Next = evenTail.Next
			node = lastOdd
			evenTail.Next = nil
		} else {
			lastOdd = node
		}
		node = node.Next
		if node == nil {
			break
		}
	}
	lastOdd.Next = evenHead
	return head
}

func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	fmt.Println(h)
	fmt.Println(oddEvenList(h))

	hh := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Println(hh)
	fmt.Println(oddEvenList(hh))
}
