package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric1(root.Left, root.Right)
}

func isSymmetric1(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isSymmetric1(left.Left, right.Right) && isSymmetric1(left.Right, right.Left)
}

func main() {
	// tree := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	tree := &TreeNode{9, &TreeNode{-42, nil, &TreeNode{76, &TreeNode{13, nil, nil}, nil}}, &TreeNode{-42, &TreeNode{76, &TreeNode{13, nil, nil}, nil}, nil}}
	fmt.Println(isSymmetric(tree))
}
