package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1 = reverseList(l1)
	l2 = reverseList(l2)
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
	return reverseList(head)
}

func reverseList(head *ListNode) *ListNode {
	tail := head
	for ; tail != nil && tail.Next != nil; tail = tail.Next {

	}
	head, _ = doReverse(head, tail)
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
