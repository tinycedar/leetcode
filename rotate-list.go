package main

import (
	"fmt"
	"bytes"
)

type listNode struct {
	Val  int
	Next *listNode
}

func (n *listNode) String() string {
	var buf bytes.Buffer
	for nn := n; nn != nil; nn = nn.Next {
		buf.WriteString(fmt.Sprintf("[%d]", nn.Val))
	}
	return buf.String()
}

func rotateRight(head *listNode, k int) *listNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}
	slow, fast := head, head
	length := 1
	for ; fast.Next != nil; length++ {
		fast = fast.Next
	}
	k = k % length
	if k == 0 {
		return head
	}
	for i := 0; i < length-k-1; i++ {
		slow = slow.Next
	}
	newHead := slow.Next
	slow.Next, fast.Next = nil, head
	return newHead
}

func main() {
	head := &listNode{1, &listNode{2, &listNode{3, &listNode{4, &listNode{5, nil}}}}}
	head2 := &listNode{1, &listNode{2, nil}}
	fmt.Println(rotateRight(head, 2))
	fmt.Println(rotateRight(head2, 2))
}
