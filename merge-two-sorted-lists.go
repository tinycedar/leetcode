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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	pre1, cur1 := l1, l1
	for cur2 := l2; l2 != nil; cur2 = l2 {
		if l2.Val <= cur1.Val { // insert prior to cur1
			l2 = l2.Next
			cur2.Next = cur1
			if pre1 != cur1 {
				pre1.Next = cur2
			} else {
				l1, pre1 = cur2, cur2
			}
			cur1 = cur2
		} else { // insert after cur1
			if cur1.Next != nil && l2.Val > cur1.Next.Val {
				pre1 = pre1.Next
				cur1 = cur1.Next
				continue
			} else {
				l2 = l2.Next
				cur2.Next = cur1.Next
				cur1.Next = cur2

				// 以下代码只能在l2升序的前提下
				pre1 = cur1
				cur1 = cur2
			}
		}
	}
	return l1
}

func main() {
	fmt.Println("final: ", mergeTwoLists(&ListNode{-9, &ListNode{3, nil}}, &ListNode{5, &ListNode{7, nil}}))
	fmt.Println("final: ", mergeTwoLists(&ListNode{5, nil}, &ListNode{1, &ListNode{2, &ListNode{4, nil}}}))
	fmt.Println("final: ", mergeTwoLists(&ListNode{5, nil}, &ListNode{7, &ListNode{2, &ListNode{9, nil}}}))
	fmt.Println("final: ", mergeTwoLists(&ListNode{2, nil}, &ListNode{1, nil}))
	fmt.Println("final: ", mergeTwoLists(&ListNode{1, nil}, &ListNode{2, nil}))
}
