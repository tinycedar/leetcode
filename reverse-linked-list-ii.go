package main

import (
	"fmt"
	"bytes"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	var buf bytes.Buffer
	for nn := n; nn != nil; nn = nn.Next {
		buf.WriteString(fmt.Sprintf("[%d]", nn.Val))
	}
	return buf.String()
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil || m <= 0 || n <= 0 || m >= n {
		return head
	}
	mNode, nNode := findSubList(head, m, n)
	if mNode == nil || nNode == nil || mNode == nNode {
		return head
	}
	nextSubNode := nNode.Next
	h, t := doReverse(mNode, nNode)
	t.Next = nextSubNode
	if t == head {
		return h
	}

	for n := head; n != nil; n = n.Next {
		if n.Next == t {
			n.Next = h
			break
		}
	}
	return head
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

func findSubList(head *ListNode, m, n int) (mNode, nNode *ListNode) {
	if head == nil {
		return nil, nil
	}
	mNode, nNode = head, head
	for i := 1; i < n; i++ {
		if i < m {
			mNode = mNode.Next
			if mNode == nil {
				return
			}
		}
		nNode = nNode.Next
		if nNode == nil {
			return
		}
	}
	return
}

func main() {
	h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	//h := &ListNode{3, &ListNode{5, nil}}
	fmt.Println(h)
	fmt.Println(reverseBetween(h, 3, 4))
}
