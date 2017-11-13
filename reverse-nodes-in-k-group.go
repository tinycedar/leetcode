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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	length := 0
	for n := head; n != nil; n = n.Next {
		length++
	}
	if k > length {
		return head
	}
	if k == length {
		tail := head
		for ; tail != nil && tail.Next != nil; tail = tail.Next {
		}
		h, _ := doReverse(head, tail)
		return h
	}
	group := (length-1)/k + 1
	groupList := make([]struct {
		h *ListNode
		t *ListNode
	}, group)

	node := head
	for g := 0; g < group; g++ {
		var pair = struct {
			h *ListNode
			t *ListNode
		}{node, nil}
		for i := 0; node != nil; node = node.Next {
			if i >= k-1 {
				pair.t = node
				break
			}
			i++
		}
		if node != nil {
			node = node.Next
		}
		if pair.t != nil {
			pair.t.Next = nil
		}
		groupList[g] = pair
	}
	for i := range groupList {
		gg := groupList[i]
		h, t := doReverse(gg.h, gg.t)
		gg.h = h
		gg.t = t
		groupList[i] = gg
	}
	for i := 0; i < group-1; i++ {
		next := groupList[i+1]
		if next.h != nil {
			groupList[i].t.Next = next.h
		} else {
			groupList[i].t.Next = next.t
		}
	}

	return groupList[0].h
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
	fmt.Println(h)
	fmt.Println("k = ", 2, "\t", reverseKGroup(h, 2))
	fmt.Println("k = ", 3, "\t", reverseKGroup(h, 3))
}
