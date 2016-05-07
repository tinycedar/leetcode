package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := []string{}
	for _, node := range binaryPath(root) {
		path := ""
		for i, num := range node {
			if i == 0 {
				path = fmt.Sprintf("%d", num)
			} else {
				path = fmt.Sprintf("%s->%d", path, num)
			}
		}
		result = append(result, path)
	}
	return result
}

func binaryPath(node *TreeNode) [][]int {
	if node == nil {
		return nil
	}
	path := [][]int{}
	current := []int{node.Val}
	if node.Left != nil {
		leftPath := binaryPath(node.Left)
		fmt.Println("Current node: ", node.Val, "\tleftPath: ", leftPath)
		if len(leftPath) > 0 {
			for _, childPath := range leftPath {
				if len(childPath) <= 0 {
					continue
				}
				path = append(path, append(current, childPath...))
			}
		}
	}
	if node.Right != nil {
		rightPath := binaryPath(node.Right)
		fmt.Println("Current node: ", node.Val, "\trightPath: ", rightPath)
		if len(rightPath) > 0 {
			for _, childPath := range rightPath {
				if len(childPath) <= 0 {
					continue
				}
				path = append(path, append(current, childPath...))
			}
		}
	}
	if len(path) == 0 {
		path = append(path, current)
	}
	return path
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, nil}, &TreeNode{4, nil, nil}}
	fmt.Println(binaryTreePaths(root))
}
