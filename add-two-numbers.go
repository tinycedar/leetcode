package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, tail *ListNode
	jinwei := 0
	for left, right := l1, l2; left != nil || right != nil; {
		val := 0
		if left != nil {
			val += left.Val
			left = left.Next
		}
		if right != nil {
			val += right.Val
			right = right.Next
		}
		val += jinwei
		node := &ListNode{val % 10, nil}
		if head == nil {
			head, tail = node, node
		} else {
			tail.Next = node
			tail = tail.Next
		}
		jinwei = val / 10

		if (left == nil || right == nil) && jinwei > 0 {
			tail.Next = &ListNode{jinwei, nil}
		}
	}
	return head
}
