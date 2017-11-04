package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	return getMinNode(lists)
}

func getMinNode(heads []*ListNode) *ListNode {
	if heads == nil || len(heads) == 0 {
		return nil
	}
	var minNode *ListNode
	var minPos = -1
	for i, head := range heads {
		if head == nil {
			continue
		}
		if minNode == nil || minNode.Val > head.Val {
			minNode = head
			minPos = i
		}
	}
	if minNode == nil {
		return nil
	}
	heads[minPos] = heads[minPos].Next
	minNode.Next = getMinNode(heads)
	return minNode
}
