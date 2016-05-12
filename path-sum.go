package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root != nil {
		for _, p := range pathSum(root) {
			if sum == p {
				return true
			}
		}
	}
	return false
}

func pathSum(node *TreeNode) []int {
	if node == nil {
		return nil
	}
	path := []int{}
	if node.Left != nil {
		for _, p := range pathSum(node.Left) {
			path = append(path, node.Val+p)
		}
	}
	if node.Right != nil {
		for _, p := range pathSum(node.Right) {
			path = append(path, node.Val+p)
		}
	}
	if node.Left == nil && node.Right == nil {
		path = append(path, node.Val)
	}
	return path
}

func main() {
	tree := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, nil}
	fmt.Println(hasPathSum(tree, 22))
}
