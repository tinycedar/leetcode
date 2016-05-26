package main

import (
	"fmt"
)

type treeNode struct {
	Val   int
	left  *treeNode
	right *treeNode
}

func hasPathSum(root *treeNode, sum int) bool {
	if root != nil {
		for _, p := range pathSum(root) {
			if sum == p {
				return true
			}
		}
	}
	return false
}

func pathSum(node *treeNode) []int {
	if node == nil {
		return nil
	}
	path := []int{}
	if node.left != nil {
		for _, p := range pathSum(node.left) {
			path = append(path, node.Val+p)
		}
	}
	if node.right != nil {
		for _, p := range pathSum(node.right) {
			path = append(path, node.Val+p)
		}
	}
	if node.left == nil && node.right == nil {
		path = append(path, node.Val)
	}
	return path
}

func main() {
	tree := &treeNode{5, &treeNode{4, &treeNode{11, &treeNode{7, nil, nil}, &treeNode{2, nil, nil}}, nil}, nil}
	fmt.Println(hasPathSum(tree, 22))
}
