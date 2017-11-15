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

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	length := 0
	tail := head
	for n := head; n != nil; n = n.Next {
		if n.Next == nil {
			tail = n
		}
		length++
	}
	firstTail := head
	for n, i := head, 1; i <= length; i++ {
		if i == length/2 {
			firstTail = n
		}
		n = n.Next
	}
	secondHead := firstTail.Next
	firstTail.Next = nil

	secondHead, tail = doReverse(secondHead, tail)
	for left, right := head, secondHead; right != nil; {
		nextRight := right.Next
		right.Next = left.Next
		left.Next = right
		if right.Next != nil {
			left = right.Next
		} else {
			left = right
		}
		right = nextRight
	}

}

func doReverse(head, tail *ListNode) (*ListNode, *ListNode) {
	if head == nil || tail == nil || head == tail {
		return tail, head
	}
	pre, cur, next := head, head.Next, head.Next.Next
	for {
		cur.Next = pre
		if cur == tail {
			break
		}
		pre = cur
		cur = next
		next = next.Next
	}
	head.Next = nil
	return tail, head
}

func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	hh := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	fmt.Print(h)
	reorderList(h)
	fmt.Println("\t", h)

	fmt.Print(hh)
	reorderList(hh)
	fmt.Println("\t", hh)

}
